import { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import { service, externalService } from "./request"
import tool from "./tool";
import { get, isEmpty } from 'lodash';
import qs from 'qs';
import { Response } from "@/api/manage/base";

// Interface for response data (adjust based on your API response structure)
export interface ApiResponse<T = any> {
    code: number;
    message: string;
    success?: boolean;
    data?: T;
    size?: number;
}

// Interface for custom request configuration
interface CustomRequestConfig extends AxiosRequestConfig {
    headers?: Record<string, string>;
    params?: Record<string, any>;
    data?: Record<string, any>;
    url?: string;
}

// Interface for environment variables (from import.meta.env)
interface Env {
    VITE_APP_TOKEN_PREFIX: string;
    VITE_APP_ID: string;
    VITE_APP_OPEN_PROXY: string;
    VITE_APP_PROXY_PREFIX: string;
    VITE_APP_BASE_URL: string;
}


function stringify(data) {
    return qs.stringify(data, { allowDots: true, encode: false });
}

/**
 * @description Creates a request function
 * @param service Axios instance for internal requests
 * @param externalService Axios instance for external requests
 * @returns A function that makes HTTP requests with the specified configuration
 */
function createRequest<T>(
    service: AxiosInstance,
    externalService: AxiosInstance,
): (config: CustomRequestConfig) => Promise<Response<T> | ApiResponse> {
    // @ts-ignore
    return function (config: CustomRequestConfig): Promise<AxiosResponse | ApiResponse> {
        // 忽略报错
        // @ts-ignore
        const env: Env = import.meta.env;
        const token: string | null = tool.local.get(env.VITE_APP_TOKEN_PREFIX); // Assuming tool.local.get returns string | null
        // @ts-ignore
        const setting: { language?: string } | null = tool.local.get('setting'); // Assuming setting is an object
        const appId: string = env.VITE_APP_ID;

        const configDefault: CustomRequestConfig = {
            headers: {
                Authorization: `Bearer ${token}`,
                'Accept-Language': setting?.language || 'zh_CN',
                'Content-Type': get(
                    config,
                    'headers.Content-Type',
                    'application/json;charset=UTF-8',
                ),
                'X-App-Id': appId,
            },
            timeout: 60000,
            data: {},
        };

        delete config.headers;
        const option: CustomRequestConfig = { ...configDefault, ...config };

        // Handle query parameters
        if (!isEmpty(option.params)) {
            option.url = `${option.url}?${stringify(option.params)}`;
            option.params = {};
        }

        // Determine base URL based on whether the URL is external
        if (!/^(http|https)/g.test(option.url!)) {
            option.baseURL =
                env.VITE_APP_OPEN_PROXY === 'true'
                    ? env.VITE_APP_PROXY_PREFIX
                    : env.VITE_APP_BASE_URL;
            return service(option);
        } else {
            return externalService(option);
        }
    };
}
export const http: <T = any>(config: CustomRequestConfig) => Promise<Response<T> | ApiResponse<T>> = createRequest(service, externalService);