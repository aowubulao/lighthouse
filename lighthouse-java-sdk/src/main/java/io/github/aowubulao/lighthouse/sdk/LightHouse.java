package io.github.aowubulao.lighthouse.sdk;

import cn.hutool.core.codec.Base64;
import cn.hutool.cron.CronUtil;
import cn.hutool.cron.task.Task;
import cn.hutool.crypto.symmetric.AES;
import cn.hutool.http.HttpRequest;
import cn.hutool.http.HttpResponse;
import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import io.github.aowubulao.lighthouse.sdk.api.LighthouseApi;
import io.github.aowubulao.lighthouse.sdk.config.LightHouseConfig;
import lombok.AccessLevel;
import lombok.NoArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.atomic.AtomicBoolean;

/**
 * @author Neo.Zzj
 * @date 2022/8/18
 */
@Slf4j
@NoArgsConstructor(access = AccessLevel.PRIVATE)
public class LightHouse {

    private static final AtomicBoolean SWITCH = new AtomicBoolean(false);

    private static String generateToken;

    private static final Map<String, ConfigDto> MAP = new HashMap<>(32);
    private static final Map<String, Map<String, String>> CONFIG_MAP = new HashMap<>(32);

    public static void startService() {
        if (!SWITCH.compareAndSet(false, true)) {
            log.warn("Lighthouse service has already started!");
        }
        generateTokenM();
        for (String file : LightHouseConfig.getConfigFiles()) {
            setConfigToMap(file);
        }
        CronUtil.schedule("10 * * * * ?", (Task) LightHouse::listenChange);
        CronUtil.schedule("0 20 * * * * ?", (Task) LightHouse::generateTokenSchedule);
        CronUtil.setMatchSecond(true);
        CronUtil.start();
    }

    public static String getConfig(String file, String key) {
        Map<String, String> map = CONFIG_MAP.get(file);
        return map == null ? null : map.get(key);
    }

    private static void listenChange() {
        List<String> configFiles = LightHouseConfig.getConfigFiles();
        for (String configName : configFiles) {
            listenSingle(configName);
        }
    }

    private static void listenSingle(String file) {
        String api = LightHouseConfig.getServer() + String.format(LighthouseApi.CONFIG_SET_VERSION_GET, file);
        ConfigDto configDto = getResponseBody(HttpRequest.get(api));
        if (configDto == null) {
            log.error("File {} response is null!", file);
            return;
        }
        String version = MAP.get(file).getVersion();
        if (!configDto.getVersion().equals(version)) {
            log.info("File change config, new version: {}", configDto.getVersion());
            setConfigToMap(file);
        }

    }

    @SuppressWarnings("all")
    private static void setConfigToMap(String file) {
        String api = LightHouseConfig.getServer() + LighthouseApi.CONFIG_SET_GET + file;
        ConfigDto configDto = getResponseBody(HttpRequest.get(api));
        if (configDto == null) {
            return;
        }
        MAP.put(file, configDto);
        String config = configDto.getConfig();
        log.info("File config read: {}", config);
        Map map = JSONUtil.toBean(config, Map.class);
        CONFIG_MAP.put(file, map);
    }

    private static ConfigDto getResponseBody(HttpRequest request) {
        HttpResponse response = request.contentType("application/json")
                .header("web-token", generateToken)
                .execute();
        response.close();
        if (response.isOk()) {
            JSONObject jsonObject = JSONUtil.parseObj(response.body());
            String data = jsonObject.getStr("data");
            try {
                return JSONUtil.toBean(data, ConfigDto.class);
            } catch (Exception e) {
                //
            }
        }
        return null;
    }

    private static void generateTokenSchedule() {
        generateTokenM();
    }

    private static void generateTokenM() {
        AES aes = new AES("CBC", "PKCS7Padding",
                LightHouseConfig.getAes().getBytes(), LightHouseConfig.getAes().substring(0, 16).getBytes());
        String source = LightHouseConfig.getWebPassword() + "?" + System.currentTimeMillis();
        byte[] encrypt = aes.encrypt(source);
        generateToken = Base64.encode(encrypt);
    }

}
