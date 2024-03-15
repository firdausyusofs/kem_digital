import axios, { AxiosInstance, AxiosRequestConfig, InternalAxiosRequestConfig, AxiosResponse, AxiosError, AxiosRequestHeaders } from 'axios';

// Create and configure an Axios instance
const apiService: AxiosInstance = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL, // Set your API base URL
});

apiService.interceptors.request.use(
  async (config: InternalAxiosRequestConfig<any>) => {

    let headers = {
      ...config.headers,
      'Content-Type': 'application/json',
      // Authorization: session?.access_token?`Bearer ${session.access_token}`:null,
    } as AxiosRequestHeaders

    config.headers = headers;

    return config;
  },
  (error: AxiosError) => {
    console.log(error);
    // Handle request errors
    return Promise.reject(error);
  }
);

apiService.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  (error: AxiosError) => {
    // Handle response errors
    return Promise.reject(error.response?.data);
  }
);

export default apiService;
