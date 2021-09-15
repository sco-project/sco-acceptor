/**
 *  状态码配置
 */
export default {
    /**
     * HTTP状态码
     * httpStatusCode                   HTTP状态码
     * messasge               String    信息
     * timestamp              Number    时间戳
     * data                   Object    数据
     */
    httpStatusCode: {
        CONTINUE: 100,  // 继续。客户端应继续其请求
        SWITCHING_PROTOCOLS: 101,  // 切换协议。服务器根据客户端的请求切换协议。只能切换到更高级的协议，例如，切换到HTTP的新版本协议
        OK: 200,  // 请求成功。一般用于GET与POST请求
        CREATED: 201,  // 已创建。成功请求并创建了新的资源
        ACCEPTED: 202,  // 已接受。已经接受请求，但未处理完成
        NON_AUTHORITATIVE_INFORMATION: 203,  // 非授权信息。请求成功。但返回的meta信息不在原始的服务器，而是一个副本
        NO_CONTENT: 204,  // 无内容。服务器成功处理，但未返回内容。在未更新网页的情况下，可确保浏览器继续显示当前文档
        RESET_CONTENT: 205,  // 重置内容。服务器处理成功，用户终端（例如：浏览器）应重置文档视图。可通过此返回码清除浏览器的表单域
        PARTIAL_CONTENT: 206,  // 部分内容。服务器成功处理了部分GET请求
        MULTIPLE_CHOICES: 300,  // 多种选择。请求的资源可包括多个位置，相应可返回一个资源特征与地址的列表用于用户终端（例如：浏览器）选择
        MOVED_PERMANENTLY: 301,  // 永久移动。请求的资源已被永久的移动到新URI，返回信息会包括新的URI，浏览器会自动定向到新URI。今后任何新的请求都应使用新的URI代替
        FOUND: 302,  // 临时移动。与301类似。但资源只是临时被移动。客户端应继续使用原有URI
        SEE_OTHER: 303,  // 查看其它地址。与301类似。使用GET和POST请求查看
        NOT_MODIFIED: 304,  // 未修改。所请求的资源未修改，服务器返回此状态码时，不会返回任何资源。客户端通常会缓存访问过的资源，通过提供一个头信息指出客户端希望只返回在指定日期之后修改的资源
        USE_PROXY: 305,  // 使用代理。所请求的资源必须通过代理访问
        UNUSED: 306,  // 已经被废弃的HTTP状态码
        TEMPORARY_REDIRECT: 307,  // 临时重定向。与302类似。使用GET请求重定向
        BAD_REQUEST: 400,  // 客户端请求的语法错误，服务器无法理解
        UNAUTHORIZED: 401,  // 请求要求用户的身份认证
        PAYMENT_REQUIRED: 402,  // 保留，将来使用
        FORBIDDEN: 403,  // 服务器理解请求客户端的请求，但是拒绝执行此请求
        NOT_FOUND: 404,  // 服务器无法根据客户端的请求找到资源（网页）。通过此代码，网站设计人员可设置"您所请求的资源无法找到"的个性页面
        METHOD_NOT_ALLOWED: 405,  // 客户端请求中的方法被禁止
        NOT_ACCEPTABLE: 406,  // 服务器无法根据客户端请求的内容特性完成请求
        PROXY_AUTHENTICATION_REQUIRED: 407,  // 请求要求代理的身份认证，与401类似，但请求者应当使用代理进行授权
        REQUEST_TIMEOUT: 408,  // 服务器等待客户端发送的请求时间过长，超时
        CONFLICT: 409,  // 服务器完成客户端的PUT请求是可能返回此代码，服务器处理请求时发生了冲突
        GONE: 410,  // 客户端请求的资源已经不存在。410不同于404，如果资源以前有现在被永久删除了可使用410代码，网站设计人员可通过301代码指定资源的新位置
        LENGTH_REQUIRED: 411,  // 服务器无法处理客户端发送的不带Content-Length的请求信息
        PRECONDITION_FAILED: 412,  // 客户端请求信息的先决条件错误
        REQUEST_ENTITY_TOO_LARGE: 413,  // 由于请求的实体过大，服务器无法处理，因此拒绝请求。为防止客户端的连续请求，服务器可能会关闭连接。如果只是服务器暂时无法处理，则会包含一个Retry-After的响应信息
        REQUEST_URI_TOO_LARGE: 414,  // 请求的URI过长（URI通常为网址），服务器无法处理
        UNSUPPORTED_MEDIA_TYPE: 415,  // 服务器无法处理请求附带的媒体格式
        REQUESTED_RANGE_NOT_SATISFIABLE: 416,  // 客户端请求的范围无效
        EXPECTATION_FAILED: 417,  // 服务器无法满足Expect的请求头信息
        INTERNAL_SERVER_ERROR: 500,  // 服务器内部错误，无法完成请求
        NOT_IMPLEMENTED: 501,  // 服务器不支持请求的功能，无法完成请求
        BAD_GATEWAY: 502,  // 充当网关或代理的服务器，从远端服务器接收到了一个无效的请求
        SERVICE_UNAVAILABLE: 503,  // 由于超载或系统维护，服务器暂时的无法处理客户端的请求。延时的长度可包含在服务器的Retry-After头信息中
        GATEWAY_TIMEOUT: 504,  // 充当网关或代理的服务器，未及时从远端服务器获取请求
        HTTP_VERSION_NOT_SUPPORTED: 505,  // 服务器不支持请求的HTTP协议的版本，无法完成处理
    },

    /**
     * 响应状态码
     * responseStatusCode               响应状态码
     * messasge               String    信息
     * timestamp              Number    时间戳
     * data                   Object    数据
     */
    responseStatusCode: {
        SUCCESS: 0,        // 请求成功
        BUSY: -1,       // 系统繁忙

        // 过期的
        EXPIRED: 4000,     // 过期的
        EXPIRED_TOKEN: 4001,     // 过期的令牌

        // 无效的
        INVALID_ARGUMENTS: 4100,     // 无效的参数
        INVALID_SIGN: 4101,     // 无效的签名
        INVALID_UPLOAD_FILE_TYPE: 4102,     // 不合法的文件类型
        INVALID_UPLOAD_FILE_SIZE: 4103,     // 不合法的文件大小
        INVALID_PASSWORD: 4104,     // 无效的密码
        INVALID_CAPTCHA: 4105,     // 无效的验证码
        INVALID_MUCH_TIME: 4106,     // 验证码输入次数过多

        // 缺少的
        MISSING_ARGUMENTS: 4200,     // 缺少参数

        // 重复的
        REPEAT: 4300,     // 重复记录
        REPEAT_ACCOUNT: 4301,     // 账号已存在

        // 未找到的
        NOT_FOUND: 4400,     // 没有找到记录
        NOT_FOUND_USER: 4401,     // 不存在的用户
        NOT_FOUND_SERVICE: 4402,     // 不存在的服务
        NOT_FOUND_RESOURCE: 4403,     // 不存在的资源
        NOT_FOUND_FILE: 4404,     // 不存在的资源

        // 超时的
        TIMEOUT: 4500,     // 超时

        // 拒绝的
        DENIED_ACCESS: 4600,     // 拒绝访问
        DENIED_UNAUTHORIZED: 4601,     // 未授权:由于凭据无效,访问被拒绝
        DENIED_ACCOUNT_DISABLED: 4602,     // 账号禁用
        DENIED_LOGIN_TOO_MANY: 4603,     // 登录次数过多

        // 非法的操作
        ILLEGAL_OPERATE: 4700,     // 非法的操作
        ILLEGAL_ARGUMENT: 4701,     // 非法的参数

        // 不支持的
        NO_SUPPORT: 4800,     // 不支持的服务

        // 未注册的
        UNTRUSTED: 4900,     // 未注册的
        UNREGISTERED_RESOURCE: 4901,     // 未注册的资源

        // 业务错误
        SYSTEM_EXCEPTION: 5000,     // 系统异常
        BUSINESS_EXCEPTION: 5001,     // 业务异常
    },
}

