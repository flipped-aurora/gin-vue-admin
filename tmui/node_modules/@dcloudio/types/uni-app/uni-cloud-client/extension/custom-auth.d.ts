declare namespace UniCloudNamespace {
  interface UserInfo {
    /**
     * 用户在uniCloud的唯一ID
     */
    uid: string;
    /**
     * 自定义登录的用户ID
     */
    customUserId: string;
  }

  interface Auth {
    /**
     * 任何方式登录成功后，可以调用 getUserInfo 获得用户的身份信息
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/authentication?id=authgetuserinfo](https://uniapp.dcloud.io/uniCloud/authentication?id=authgetuserinfo)
     */
    getUserInfo(): UserInfo;
    /**
     * 开发者可以通过 getLoginState() 来获取当前的登录状态，调用 getLoginState() 后，SDK 会识别本地是否有登录状态，如果有，则会尝试刷新登录状态，若刷新登录状态成功，则会返回新的登录状态，否则返回 undefined
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/authentication?id=authgetloginstate](https://uniapp.dcloud.io/uniCloud/authentication?id=authgetloginstate)
     */
    getLoginState(): void;
    /**
     * 进行匿名登录
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/authentication?id=authsigninanonymously](https://uniapp.dcloud.io/uniCloud/authentication?id=authsigninanonymously)
     */
    signInAnonymously(): void;
    /**
     * 进行自定义登录
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/authentication?id=authsigninwithticket](https://uniapp.dcloud.io/uniCloud/authentication?id=authsigninwithticket)
     */
    signInWithTicket(): void;
    /**
     * 进行自定义登录
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/authentication?id=authshouldrefreshaccesstoken](https://uniapp.dcloud.io/uniCloud/authentication?id=authshouldrefreshaccesstoken)
     */
    shouldRefreshAccessToken(callback: (result: any) => void): void;
  }

  interface UniCloud {
    /**
     * 获取登录对象
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/authentication?id=custom-auth](https://uniapp.dcloud.io/uniCloud/authentication?id=custom-auth)
     */
    customAuth(): Auth;
  }
}
