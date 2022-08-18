package io.github.aowubulao.lighthouse.sdk.config;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.List;

/**
 * @author Neo.Zzj
 * @date 2022/8/18
 */
class LightHouseConfigTest {

    @Test
    void getWebPassword() {
        String webPassword = LightHouseConfig.getWebPassword();
        System.out.println(webPassword);
        Assertions.assertNotNull(webPassword);
    }

    @Test
    void getAes() {
        String get = LightHouseConfig.getAes();
        System.out.println(get);
        Assertions.assertNotNull(get);
    }

    @Test
    void getConfigFiles() {
        List<String> configFiles = LightHouseConfig.getConfigFiles();
        System.out.println(configFiles);
        Assertions.assertNotNull(configFiles);
    }
}
