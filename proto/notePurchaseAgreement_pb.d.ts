// package: notepurchaseagreement
// file: notePurchaseAgreement.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as protoc_gen_openapiv2_options_annotations_pb from "./protoc-gen-openapiv2/options/annotations_pb";

export class NotePuchaseAgreement extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getFirstname(): string;
  setFirstname(value: string): void;

  getLastname(): string;
  setLastname(value: string): void;

  getDateofbirth(): string;
  setDateofbirth(value: string): void;

  getHomeaddress(): string;
  setHomeaddress(value: string): void;

  getEmailaddress(): string;
  setEmailaddress(value: string): void;

  getPhonenumber(): string;
  setPhonenumber(value: string): void;

  getSocialsecurity(): string;
  setSocialsecurity(value: string): void;

  getFundscommitted(): number;
  setFundscommitted(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NotePuchaseAgreement.AsObject;
  static toObject(includeInstance: boolean, msg: NotePuchaseAgreement): NotePuchaseAgreement.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NotePuchaseAgreement, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NotePuchaseAgreement;
  static deserializeBinaryFromReader(message: NotePuchaseAgreement, reader: jspb.BinaryReader): NotePuchaseAgreement;
}

export namespace NotePuchaseAgreement {
  export type AsObject = {
    id: number,
    firstname: string,
    lastname: string,
    dateofbirth: string,
    homeaddress: string,
    emailaddress: string,
    phonenumber: string,
    socialsecurity: string,
    fundscommitted: number,
  }
}

export class GetNotePurchaseAgreementDocRequest extends jspb.Message {
  hasPayload(): boolean;
  clearPayload(): void;
  getPayload(): NotePuchaseAgreement | undefined;
  setPayload(value?: NotePuchaseAgreement): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetNotePurchaseAgreementDocRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetNotePurchaseAgreementDocRequest): GetNotePurchaseAgreementDocRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetNotePurchaseAgreementDocRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetNotePurchaseAgreementDocRequest;
  static deserializeBinaryFromReader(message: GetNotePurchaseAgreementDocRequest, reader: jspb.BinaryReader): GetNotePurchaseAgreementDocRequest;
}

export namespace GetNotePurchaseAgreementDocRequest {
  export type AsObject = {
    payload?: NotePuchaseAgreement.AsObject,
  }
}

export class GetNotePurchaseAgreementDocResponse extends jspb.Message {
  getFilecontent(): Uint8Array | string;
  getFilecontent_asU8(): Uint8Array;
  getFilecontent_asB64(): string;
  setFilecontent(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetNotePurchaseAgreementDocResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetNotePurchaseAgreementDocResponse): GetNotePurchaseAgreementDocResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetNotePurchaseAgreementDocResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetNotePurchaseAgreementDocResponse;
  static deserializeBinaryFromReader(message: GetNotePurchaseAgreementDocResponse, reader: jspb.BinaryReader): GetNotePurchaseAgreementDocResponse;
}

export namespace GetNotePurchaseAgreementDocResponse {
  export type AsObject = {
    filecontent: Uint8Array | string,
  }
}

