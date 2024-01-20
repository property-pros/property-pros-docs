package main

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	controllers "github.com/vireocloud/property-pros-docs/server/controllers"
	"github.com/vireocloud/property-pros-docs/server/third_party"
	propertyProsApi "github.com/vireocloud/property-pros-sdk/api/note_purchase_agreement/v1"
	statementApi "github.com/vireocloud/property-pros-sdk/api/statement/v1"
)

var (
	enableTls = flag.Bool("enable_tls", false, "Use TLS - required for HTTP2.")
)

func main() {
	// controllers.chill()
	flag.Parse()
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	if *enableTls {
		fmt.Println("TLS enabled")
		// port := "9090"
		//TODO: get a key file
		creds, err := credentials.NewServerTLSFromFile("/etc/ssl/cert.pem", "")

		if err != nil {
			grpclog.Fatalln(err)
		}

		StartServer(creds)
	} else {
		fmt.Println("TLS disabled")
		StartServer(insecure.NewCredentials())
	}
}

func grpcHandlerFunc(grpcServer *grpc.Server, grpcWebServer *grpcweb.WrappedGrpcServer, restHandler http.Handler, oa http.Handler) http.Handler {

	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling request")
		grpclog.Infof("url: %v\r\n\r\ncontent type: %v\r\n\r\n", r.URL.Path, r.Header.Get("Content-Type"))
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc-web+proto") {
			grpclog.Infoln("grpc-web request")
			grpcWebServer.ServeHTTP(w, r)
			return
		}

		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpclog.Infoln("grpc request")
			grpcServer.ServeHTTP(w, r)
			return
		}

		if matched, err := regexp.MatchString("v\\d", r.URL.Path); err == nil && matched {
			grpclog.Infoln("rest api request")
			restHandler.ServeHTTP(w, r)
			return
		}

		oa.ServeHTTP(w, r)

	}), &http2.Server{})
}

func StartServer(transportCredentials credentials.TransportCredentials) {
	wg := sync.WaitGroup{}

	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*1024),
		grpc.MaxSendMsgSize(1024*1024*1024),
	)

	wrappedServer := grpcweb.WrapServer(grpcServer)

	propertyProsApi.RegisterNotePurchaseAgreementServiceServer(grpcServer, &controllers.NotePurchaseAgreementController{})
	statementApi.RegisterStatementServiceServer(grpcServer, &controllers.StatementController{})

	gwmux := runtime.NewServeMux()

	ctx := context.Background()
	// server not accessible from property-pros-service, but works with postman
	dopts := []grpc.DialOption{
		grpc.WithTransportCredentials(transportCredentials),
	}

	host := "0.0.0.0"
	serverPort := "8020"
	gatewayPort := "8023"
	scheme := "dns:///"

	serverUrl := fmt.Sprintf("%v:%v", host, serverPort)
	dialUrl := fmt.Sprintf("%vlocalhost:%v", scheme, gatewayPort)

	fmt.Println("server url: ", serverUrl)
	fmt.Println("dial url: ", dialUrl)

	wg.Add(1)

	go func() {
		// Create a server
		srv := &http.Server{
			Addr:    serverUrl,
			Handler: grpcHandlerFunc(grpcServer, wrappedServer, gwmux, getOpenAPIHandler()),
		}

		// Configure the server with the http2 options
		http2Server := http2.Server{
			MaxReadFrameSize: 1024 * 1024 * 10, // maximum allowed frame size
		}

		_ = http2.ConfigureServer(srv, &http2Server)
		fmt.Printf("Listening on 8020 with framesize %v \n", http2Server.MaxReadFrameSize)
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("Http listener failed: ", err)
			wg.Done()
		}
	}()

	err := propertyProsApi.RegisterNotePurchaseAgreementServiceHandlerFromEndpoint(ctx, gwmux, dialUrl, dopts)

	if err != nil {
		grpclog.Errorf("failed starting http server: %v", err)
	}

	os.Unsetenv("HTTP_PROXY")
	os.Unsetenv("HTTPS_PROXY")
	os.Unsetenv("NO_PROXY")
	wg.Wait()
}

func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}
