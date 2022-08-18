package io.github.aowubulao.lighthouse.sdk;

import org.junit.jupiter.api.Test;

/**
 * @author Neo.Zzj
 * @date 2022/8/18
 */
class LightHouseTest {

    @Test
    void startService() {
        LightHouse.startService();
        String config = LightHouse.getConfig("test_file", "key");
        System.out.println(config);
    }
}
