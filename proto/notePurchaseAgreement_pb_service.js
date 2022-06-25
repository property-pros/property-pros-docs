// package: notepurchaseagreement
// file: notePurchaseAgreement.proto

var notePurchaseAgreement_pb = require("./notePurchaseAgreement_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var NotePurchaseAgreementService = (function () {
  function NotePurchaseAgreementService() {}
  NotePurchaseAgreementService.serviceName = "notepurchaseagreement.NotePurchaseAgreementService";
  return NotePurchaseAgreementService;
}());

NotePurchaseAgreementService.GetNotePurchaseAgreementDoc = {
  methodName: "GetNotePurchaseAgreementDoc",
  service: NotePurchaseAgreementService,
  requestStream: false,
  responseStream: false,
  requestType: notePurchaseAgreement_pb.GetNotePurchaseAgreementDocRequest,
  responseType: notePurchaseAgreement_pb.GetNotePurchaseAgreementDocResponse
};

exports.NotePurchaseAgreementService = NotePurchaseAgreementService;

function NotePurchaseAgreementServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

NotePurchaseAgreementServiceClient.prototype.getNotePurchaseAgreementDoc = function getNotePurchaseAgreementDoc(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(NotePurchaseAgreementService.GetNotePurchaseAgreementDoc, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.NotePurchaseAgreementServiceClient = NotePurchaseAgreementServiceClient;

