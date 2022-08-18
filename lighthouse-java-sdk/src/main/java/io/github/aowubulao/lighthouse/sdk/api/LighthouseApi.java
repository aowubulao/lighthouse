package io.github.aowubulao.lighthouse.sdk.api;

import lombok.AccessLevel;
import lombok.NoArgsConstructor;

/**
 * @author Neo.Zzj
 * @date 2022/8/18
 */
@NoArgsConstructor(access = AccessLevel.PRIVATE)
public final class LighthouseApi {

    public static final String CONFIG_SET_GET = "/api/v1/configuration/";

    public static final String CONFIG_SET_VERSION_GET = "/api/v1/configuration/%s/version";

    public static final String CONFIG_PUT = "/api/v1/configuration/";

}
