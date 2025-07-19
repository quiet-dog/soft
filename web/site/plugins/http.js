// plugins/auth-fetch.js
import CryptoJS from 'crypto-js'

export default defineNuxtPlugin((nuxtApp) => {
  const tokenCookie = useCookie('auth:token')
  const expireCookie = useCookie('auth:token_expire')
  const langCookie = useCookie('language')
  const runtimeConfig = useRuntimeConfig()
  function getConfig() {
    return {
      baseURL: runtimeConfig.public.baseURL,
      appId: runtimeConfig.public.appId,
      appSecret: runtimeConfig.public.appSecret,
      defaultLang: 'zh_CN',
    }
  }

  function generateSignature(appSecret) {
    const xtimestamp = Date.now().toString()
    const xnonce = Math.floor(Math.random() * 999999999 + 99999).toString()
    const md5Func = CryptoJS.MD5
    const xsign = md5Func(appSecret + xtimestamp + xnonce).toString()
    return { timestamp: xtimestamp, nonce: xnonce, sign: xsign }
  }

  async function getToken() {
    const config = getConfig()
    if (tokenCookie.value && expireCookie.value) {
      const expireTime = Number.parseInt(expireCookie.value)
      if (Date.now() < expireTime - 60000)
        return tokenCookie.value
      return await refreshToken(tokenCookie.value)
    }
    const signatureParams = await generateSignature(config.appSecret)
    if (!signatureParams) {
      if (useHelper.isClient()) {
        const Arco = await import('@arco-design/web-vue')
        Arco.Message.error('签名生成失败')
        return null
      }
    }
    const language = langCookie.value || config.defaultLang
    const response = await $fetch('/api/getToken', {
      baseURL: config.baseURL,
      method: 'GET',
      params: {
        app_id: config.appId,
        signature: signatureParams.sign,
        timestamp: signatureParams.timestamp,
        nonce: signatureParams.nonce,
        language,
      },
    })

    if (response?.code === 0 && response.data?.token) {
      tokenCookie.value = response.data.token
      expireCookie.value = response.data.expire * 1000
      return response.data.token
    }
    return null
  }

  async function refreshToken(currentToken) {
    const config = getConfig()
    const language = langCookie.value || config.defaultLang
    const response = await $fetch('/api/refreshToken', {
      baseURL: config.baseURL,
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${currentToken}`,
        'Accept-Language': language,
        'X-App-Id': config.appId,
      },
    })

    if (response?.code === 0 && response.data?.token) {
      tokenCookie.value = response.data.token
      expireCookie.value = response.data.expire * 1000
      return response.data.token
    }
    return currentToken
  }

  async function applyOptions(options = {}) {
    const config = getConfig()
    options.baseURL = options.baseURL ?? config.baseURL
    options.initialCache = options.initialCache ?? false
    options.headers = options.headers || {}
    options.method = options.method || 'GET'
    options.timeout = 3000

    // 确保await getMyCookie在正确的上下文中调用
    const language = langCookie.value || config.defaultLang

    let headers = {
      'accept': 'application/json',
      'Accept-Language': language,
      'X-App-Id': config.appId,
    }

    const token = tokenCookie.value
    if (!token) {
      const newToken = await getToken()
      if (newToken)
        headers.Authorization = `Bearer ${newToken}`
    }
    else {
      headers.Authorization = `Bearer ${token}`
    }

    if (useHelper.isServer()) {
      const serverHeaders = useRequestHeaders(['referer', 'cookie'])
      headers = { ...headers, ...serverHeaders }
    }

    options.headers = { ...headers, ...options.headers }
    return options
  }

  function handleError(response) {
    // 带上下文的错误显示
    const showError = async (text) => {
      if (useHelper.isClient()) {
        const Arco = await import('@arco-design/web-vue')
        Arco.Message.error(text || '未知错误')
      }
      else {
        console.error('服务端错误:', text)
      }
    }

    switch (response?.code) {
      case 1000:
        tokenCookie.value = null
        expireCookie.value = null
        showError('登录状态过期')
        break
      case 65:
        showError('资源不存在')
        break
      case 50:
        showError('服务器错误')
        break
      case 1002:
        tokenCookie.value = null
        expireCookie.value = null
        showError('token已过期')
        break
      case 1001:
      case 61:
        showError('无访问权限')
        break
      default:
        showError(response?.message || '未知错误')
    }
  }

  async function fetch(url, options = {}) {
    if (options.$) {
      const response = await fetch$(url, options)
      return {
        data: ref(response),
        pending: ref(false),
        error: ref(null),
        refresh: async () => {
          const newResponse = await fetch$(url, options)
          return newResponse
        },
      }
    }
    else {
      // 检查组件是否已挂载
      const nuxtApp = useNuxtApp()
      const isMounted = nuxtApp.isHydrating === false

      // 组件已挂载，直接使用 $fetch
      if (isMounted) {
        const response = await fetch$(url, options)
        return {
          data: ref(response),
          pending: ref(false),
          error: ref(null),
          refresh: async () => {
            const newResponse = await $fetch(url, options)
            return newResponse
          },
        }
      }
      // 组件未挂载，使用 useAsyncData
      else {
        const key = options?.key || `fetch_${url.replace(/[^a-z0-9]/gi, '_')}`
        options = await applyOptions({ ...options })
        return await useAsyncData(key, () => $fetch(url, options))
      }
    }
  }

  async function fetch$(url, options) {
    options = await applyOptions({ ...options })
    const response = await $fetch(url, options)
    return response
  }

  const http = () => {
    return {
      get: (url, params, options) => fetch(url, { method: 'GET', params, ...options }),
      post: (url, body, options) => fetch(url, { method: 'POST', body, ...options }),
      put: (url, body, options) => fetch(url, { method: 'PUT', body, ...options }),
      delete: (url, params, options) => fetch(url, { method: 'DELETE', params, ...options }),
      get$: (url, params, options) => fetch$(url, { method: 'GET', params, ...options }),
      post$: (url, body, options) => fetch$(url, { method: 'POST', body, ...options }),
      put$: (url, body, options) => fetch$(url, { method: 'PUT', body, ...options }),
      delete$: (url, params, options) => fetch$(url, { method: 'DELETE', params, ...options }),
      applyOptions: options => applyOptions(options),
    }
  }

  // 注入全局实例
  nuxtApp.provide('http', http)
})
