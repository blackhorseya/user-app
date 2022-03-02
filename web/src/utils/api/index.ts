import {ApiError, IResponseData} from '@/types';

const {MODE, VITE_API_ENDPOINT} = import.meta.env;

export var _token: string;

const _url = (url: string) =>
    MODE === 'development' || MODE === 'test'
        ? `/api/v1${url}`
        : new URL(`/api/v1${url}`, VITE_API_ENDPOINT).href;
const _authorization = (token: string) => `Bearer ${token}`;

interface IFetchParams {
  url: string;
  token?: string;
  headers?: HeadersInit;
  body?: BodyInit;
}

export async function get<T = any>({url, token, headers = {}}: IFetchParams) {
  if (token)
    headers = {
      ...headers,
      Authorization: _authorization(token),
    };
  const res = await fetch(_url(url), {
    method: 'GET',
    headers,
  });
  const json = (await res.json()) as IResponseData<T>;
  if (res.ok) return json;
  else {
    throw new ApiError<T>(res.status, json);
  }
}

export async function post<T = any>({
                                      url,
                                      body = JSON.stringify({}),
                                      token,
                                      headers = {},
                                    }: IFetchParams) {
  if (token) headers = {...headers, Authorization: _authorization(token)};
  const res = await fetch(_url(url), {
    method: 'POST',
    headers,
    body: body,
  });
  const json = (await res.json()) as IResponseData<T>;
  if (res.ok) return json;
  else {
    throw new ApiError<T>(res.status, json);
  }
}

export async function remove<T = any>({
                                        url,
                                        body = JSON.stringify({}),
                                        token,
                                        headers = {},
                                      }: IFetchParams) {
  if (token) headers = {...headers, Authorization: _authorization(token)};
  const res = await fetch(_url(url), {
    method: 'DELETE',
    headers,
    body: body,
  });
  const json = (await res.json()) as IResponseData<T>;
  if (res.ok) return json;
  else {
    throw new ApiError<T>(res.status, json);
  }
}

export async function getLoginUrl() {
  try {
    const res = await fetch(_url('/auth/login'), {
      method: 'GET',
    });
    const json = await res.json();
    if (res.ok) return json.data;
  } catch (error) {
    console.warn(error);
  }
}

export async function getProfile(token: string) {
  try {
    _token = token;
    const res = await fetch(_url('/auth/me'), {
      method: 'GET',
      headers: {
        Authorization: _authorization(token),
      },
    });
    const json = await res.json();
    if (res.ok) return json;
  } catch (error) {
    console.warn(error);
  }
}