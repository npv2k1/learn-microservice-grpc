/* eslint-disable */
import { Metadata } from "@grpc/grpc-js";
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";

export const protobufPackage = "hero";

/** hero/hero.proto */

export interface HeroById {
  id: number;
}

export interface Hero {
  id: number;
  name: string;
}

export interface ServerStreamRequest {
  num: number;
}

export interface ServerStreamResponse {
  num: number;
}

export interface ClientStreamRequest {
  num: number;
}

export interface ClientStreamResponse {
  num: number;
}

export interface TwoWayStreamRequest {
  num: number;
}

export interface TwoWayStreamResponse {
  num: number;
}

export const HERO_PACKAGE_NAME = "hero";

export interface HeroesServiceClient {
  findOne(request: HeroById, metadata: Metadata, ...rest: any): Observable<Hero>;

  serverStream(request: ServerStreamRequest, metadata: Metadata, ...rest: any): Observable<ServerStreamResponse>;

  clientStream(
    request: Observable<ClientStreamRequest>,
    metadata: Metadata,
    ...rest: any
  ): Observable<ClientStreamResponse>;

  twoWayStream(
    request: Observable<TwoWayStreamRequest>,
    metadata: Metadata,
    ...rest: any
  ): Observable<TwoWayStreamResponse>;
}

export interface HeroesServiceController {
  findOne(request: HeroById, metadata: Metadata, ...rest: any): Promise<Hero> | Observable<Hero> | Hero;

  serverStream(request: ServerStreamRequest, metadata: Metadata, ...rest: any): Observable<ServerStreamResponse>;

  clientStream(
    request: Observable<ClientStreamRequest>,
    metadata: Metadata,
    ...rest: any
  ): Promise<ClientStreamResponse> | Observable<ClientStreamResponse> | ClientStreamResponse;

  twoWayStream(
    request: Observable<TwoWayStreamRequest>,
    metadata: Metadata,
    ...rest: any
  ): Observable<TwoWayStreamResponse>;
}

export function HeroesServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ["findOne", "serverStream"];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("HeroesService", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = ["clientStream", "twoWayStream"];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("HeroesService", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const HEROES_SERVICE_NAME = "HeroesService";
