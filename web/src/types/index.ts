export interface IResponseData<T> {
  code: number;
  data?: T;
  msg: string;
}

export enum ApiStatus {
  Ok = 200,
  NotFound = 404,
  BadRequest = 400,
  ServerError = 500,
}

export class ApiError<T> extends Error {
  readonly status: ApiStatus;
  readonly data: IResponseData<T>;

  constructor(status: ApiStatus, data: IResponseData<T>) {
    super(`${status} ${data.msg}`);
    this.status = status;
    this.data = data;
  }
}