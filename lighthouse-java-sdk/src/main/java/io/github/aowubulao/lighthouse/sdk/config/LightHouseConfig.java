package io.github.aowubulao.lighthouse.sdk.config;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import java.io.IOException;
import java.io.InputStream;
import java.util.Arrays;
import java.util.List;
import java.util.Properties;

/**
 * @author Neo.Zzj
 * @date 2022/8/18
 */
@Slf4j
@NoArgsConstructor(access = AccessLevel.PRIVATE)
public class LightHouseConfig {
    @Getter
    private static String webPassword;
    @Getter
    private static String aes;
    @Getter
    private static List<String> configFiles;
    @Getter
    private static String server;

    static {
        initConfig();
    }

    private static void initConfig() {
        Properties props = new Properties();
        InputStream is = Thread.currentThread().getContextClassLoader().getResourceAsStream("lighthouse.properties");
        try {
            props.load(is);
            webPassword = props.getProperty("webPassword");
            aes = props.getProperty("aes");
            configFiles = Arrays.asList(props.getProperty("configFiles").split(","));
            server = props.getProperty("server");
        } catch (IOException e) {
            log.error("读取lighthouse配置文件出错：", e);
        }
    }

}
