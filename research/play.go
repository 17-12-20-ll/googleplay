package main

import (
   "bytes"
   "fmt"
   "github.com/89z/format/protobuf"
   "github.com/89z/googleplay"
   "io"
   "net/http"
   "net/url"
)

var body = protobuf.Message{protobuf.Tag{Number:12, Name:"string"}:"America/Mexico_City",
protobuf.Tag{Number:14, Name:"varint"}:uint64(3),
protobuf.Tag{Number:18, Name:"message"}:protobuf.Message{protobuf.Tag{Number:9, Name:"string"}:[]string{"android.ext.services", "android.ext.shared", "android.test.base", "android.test.mock", "android.test.runner", "com.android.future.usb.accessory", "com.android.ims.rcsmanager", "com.android.location.provider", "com.android.media.remotedisplay", "com.android.mediadrm.signer", "com.google.android.camera.experimental2018", "com.google.android.dialer.support", "com.google.android.gms", "com.google.android.hardwareinfo", "com.google.android.lowpowermonitordevicefactory", "com.google.android.lowpowermonitordeviceinterface", "com.google.android.maps", "com.google.android.poweranomalydatafactory", "com.google.android.poweranomalydatamodeminterface", "com.qti.snapdragon.sdk.display", "com.qualcomm.embmslibrary", "com.qualcomm.qcrilhook", "com.qualcomm.qti.QtiTelephonyServicelibrary", "com.qualcomm.qti.imscmservice@1.0-java", "com.qualcomm.qti.lpa.uimlpalibrary", "com.qualcomm.qti.ltedirectdiscoverylibrary", "com.qualcomm.qti.remoteSimlock.uimremotesimlocklibrary", "com.qualcomm.qti.uim.uimservicelibrary", "com.quicinc.cne", "com.quicinc.cneapiclient", "com.verizon.embms", "com.verizon.provider", "com.vzw.apnlib", "javax.obex", "org.apache.http.legacy"},
protobuf.Tag{Number:11, Name:"message"}:protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:uint64(7005479156896394610)},
protobuf.Tag{Number:26, Name:"message"}:[]protobuf.Message{protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.audio.low_latency",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.audio.output",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.audio.pro",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.bluetooth",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.bluetooth_le",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.hardware.camera"}, protobuf.Message{protobuf.Tag{Number:1, Name:"message"}:protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:[]uint64{0x682e64696f72646e, 0x632e657261776472, 0x796e612e6172656d}},
protobuf.Tag{Number:1, Name:"string"}:"android.hardware.camera.any",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.camera.autofocus",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.camera.capability.manual_post_processing",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.hardware.camera.capability.manual_sensor"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.camera.capability.raw",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.camera.flash",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.hardware.camera.front"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.camera.level.full",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.faketouch",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.fingerprint",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.location",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.location.gps",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.location.network",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.microphone",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"message"}:protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:[]uint64{0x682e64696f72646e, 0x6d2e657261776472},
protobuf.Tag{Number:13, Name:"fixed64"}:uint64(7308901739622527587)}}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.nfc",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.nfc.any",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.nfc.hce",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.nfc.hcef",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.opengles.aep",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"message"}:protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:[]uint64{0x682e64696f72646e, 0x722e657261776472, 0x6c616d726f6e2e6d}},
protobuf.Tag{Number:1, Name:"string"}:"android.hardware.ram.normal"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.screen.landscape",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.screen.portrait",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.accelerometer",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.assist",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.barometer",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.compass",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.gyroscope",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.hifi_sensors",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.light",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.proximity"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.stepcounter",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.sensor.stepdetector",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.strongbox_keystore",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.telephony",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.telephony.carrierlock",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.telephony.cdma",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.telephony.euicc",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.telephony.gsm",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.touchscreen",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.hardware.touchscreen.multitouch"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.touchscreen.multitouch.distinct",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.touchscreen.multitouch.jazzhand",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.usb.accessory",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.usb.host",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.vulkan.compute",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.vulkan.level",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.vulkan.version",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.wifi",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"message"}:protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:[]uint64{0x682e64696f72646e, 0x772e657261776472},
protobuf.Tag{Number:13, Name:"fixed64"}:uint64(7310012310535170406)},
protobuf.Tag{Number:1, Name:"string"}:"android.hardware.wifi.aware",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.hardware.wifi.direct"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.wifi.passpoint",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.hardware.wifi.rtt",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.activities_on_secondary_displays",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.app_widgets",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.software.autofill"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.backup",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.cant_save_state",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.companion_device_setup",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.connectionservice",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.cts",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.device_admin",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.device_id_attestation",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.file_based_encryption",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.home_screen",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.input_methods",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.live_wallpaper",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.software.managed_users"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.midi",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.picture_in_picture",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.software.print"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.securely_removes_users",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.sip",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.sip.voip",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.software.verified_boot"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"android.software.voice_recognizers",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"android.software.webview"}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.apps.dialer.SUPPORTED",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.apps.photos.PIXEL_2019_MIDYEAR_PRELOAD",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.EXCHANGE_6_2",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.GOOGLE_BUILD",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.GOOGLE_EXPERIENCE",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.PIXEL_2017_EXPERIENCE",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.PIXEL_2018_EXPERIENCE",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.PIXEL_2019_MIDYEAR_EXPERIENCE",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.PIXEL_EXPERIENCE",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.TURBO_PRELOAD",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.WELLBEING",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.android.feature.ZERO_TOUCH",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.google.hardware.camera.easel_2018",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.verizon.hardware.telephony.ehrpd",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}, protobuf.Message{protobuf.Tag{Number:1, Name:"string"}:"com.verizon.hardware.telephony.lte",
protobuf.Tag{Number:2, Name:"varint"}:uint64(0)}},
protobuf.Tag{Number:10, Name:"string"}:[]string{"android.hardware.audio.low_latency", "android.hardware.audio.output", "android.hardware.audio.pro", "android.hardware.bluetooth", "android.hardware.bluetooth_le", "android.hardware.camera", "android.hardware.camera.any", "android.hardware.camera.autofocus", "android.hardware.camera.capability.manual_post_processing", "android.hardware.camera.capability.manual_sensor", "android.hardware.camera.capability.raw", "android.hardware.camera.flash", "android.hardware.camera.front", "android.hardware.camera.level.full", "android.hardware.faketouch", "android.hardware.fingerprint", "android.hardware.location", "android.hardware.location.gps", "android.hardware.location.network", "android.hardware.microphone", "android.hardware.nfc", "android.hardware.nfc.any", "android.hardware.nfc.hce", "android.hardware.nfc.hcef", "android.hardware.opengles.aep", "android.hardware.ram.normal", "android.hardware.screen.landscape", "android.hardware.screen.portrait", "android.hardware.sensor.accelerometer", "android.hardware.sensor.assist", "android.hardware.sensor.barometer", "android.hardware.sensor.compass", "android.hardware.sensor.gyroscope", "android.hardware.sensor.hifi_sensors", "android.hardware.sensor.light", "android.hardware.sensor.proximity", "android.hardware.sensor.stepcounter", "android.hardware.sensor.stepdetector", "android.hardware.strongbox_keystore", "android.hardware.telephony", "android.hardware.telephony.carrierlock", "android.hardware.telephony.cdma", "android.hardware.telephony.euicc", "android.hardware.telephony.gsm", "android.hardware.touchscreen", "android.hardware.touchscreen.multitouch", "android.hardware.touchscreen.multitouch.distinct", "android.hardware.touchscreen.multitouch.jazzhand", "android.hardware.usb.accessory", "android.hardware.usb.host", "android.hardware.vulkan.compute", "android.hardware.vulkan.level", "android.hardware.vulkan.version", "android.hardware.wifi", "android.hardware.wifi.aware", "android.hardware.wifi.direct", "android.hardware.wifi.passpoint", "android.hardware.wifi.rtt", "android.software.activities_on_secondary_displays", "android.software.app_widgets", "android.software.autofill", "android.software.backup", "android.software.cant_save_state", "android.software.companion_device_setup", "android.software.connectionservice", "android.software.cts", "android.software.device_admin", "android.software.device_id_attestation", "android.software.file_based_encryption", "android.software.home_screen", "android.software.input_methods", "android.software.live_wallpaper", "android.software.managed_users", "android.software.midi", "android.software.picture_in_picture", "android.software.print", "android.software.securely_removes_users", "android.software.sip", "android.software.sip.voip", "android.software.verified_boot", "android.software.voice_recognizers", "android.software.webview", "com.google.android.apps.dialer.SUPPORTED", "com.google.android.apps.photos.PIXEL_2019_MIDYEAR_PRELOAD", "com.google.android.feature.EXCHANGE_6_2", "com.google.android.feature.GOOGLE_BUILD", "com.google.android.feature.GOOGLE_EXPERIENCE", "com.google.android.feature.PIXEL_2017_EXPERIENCE", "com.google.android.feature.PIXEL_2018_EXPERIENCE", "com.google.android.feature.PIXEL_2019_MIDYEAR_EXPERIENCE", "com.google.android.feature.PIXEL_EXPERIENCE", "com.google.android.feature.TURBO_PRELOAD", "com.google.android.feature.WELLBEING", "com.google.android.feature.ZERO_TOUCH", "com.google.hardware.camera.easel_2018", "com.verizon.hardware.telephony.ehrpd", "com.verizon.hardware.telephony.lte"},
protobuf.Tag{Number:12, Name:"varint"}:uint64(1080),
protobuf.Tag{Number:1, Name:"varint"}:uint64(3),
protobuf.Tag{Number:4, Name:"varint"}:uint64(2),
protobuf.Tag{Number:5, Name:"varint"}:uint64(0),
protobuf.Tag{Number:6, Name:"varint"}:uint64(0),
protobuf.Tag{Number:7, Name:"varint"}:uint64(490),
protobuf.Tag{Number:8, Name:"varint"}:uint64(196610),
protobuf.Tag{Number:14, Name:"string"}:[]string{"af", "af_ZA", "am", "am_ET", "ar", "ar_EG", "ar_SA", "ar_XB", "as", "ast", "az", "be", "be_BY", "bg", "bg_BG", "bh_IN", "bn", "bs", "ca", "ca_ES", "cs", "cs_CZ", "cy_GB", "da", "da_DK", "de", "de_DE", "el", "el_GR", "en", "en_AU", "en_CA", "en_GB", "en_IN", "en_US", "en_XA", "es", "es_ES", "es_US", "et", "et_EE", "eu", "fa", "fa_IR", "fi", "fi_FI", "fil", "fil_PH", "fr", "fr_CA", "fr_FR", "gl", "gl_ES", "gu", "hi", "hi_IN", "hr", "hr_HR", "hu", "hu_HU", "hy", "in", "in_ID", "is", "it", "it_IT", "iw", "iw_IL", "ja", "ja_JP", "ka", "kab_DZ", "kk", "km", "kn", "ko", "ko_KR", "ky", "lo", "lt", "lt_LT", "lv", "lv_LV", "mk", "ml", "mn", "mr", "ms", "ms_MY", "my", "nb", "nb_NO", "ne", "nl", "nl_NL", "or", "pa", "pa_IN", "pl", "pl_PL", "pt", "pt_BR", "pt_PT", "ro", "ro_RO", "ru", "ru_RU", "sc_IT", "si", "sk", "sk_SK", "sl", "sl_SI", "sq", "sr", "sr_Latn", "sr_RS", "sv", "sv_SE", "sw", "sw_TZ", "ta", "te", "th", "th_TH", "tr", "tr_TR", "uk", "uk_UA", "ur", "uz", "vi", "vi_VN", "zh_CN", "zh_HK", "zh_TW", "zu", "zu_ZA"},
protobuf.Tag{Number:15, Name:"string"}:[]string{"GL_AMD_compressed_ATC_texture", "GL_AMD_performance_monitor", "GL_ANDROID_extension_pack_es31a", "GL_APPLE_texture_2D_limited_npot", "GL_ARB_vertex_buffer_object", "GL_ARM_shader_framebuffer_fetch_depth_stencil", "GL_EXT_EGL_image_array", "GL_EXT_EGL_image_external_wrap_modes", "GL_EXT_EGL_image_storage", "GL_EXT_YUV_target", "GL_EXT_blend_func_extended", "GL_EXT_blit_framebuffer_params", "GL_EXT_buffer_storage", "GL_EXT_clip_control", "GL_EXT_clip_cull_distance", "GL_EXT_color_buffer_float", "GL_EXT_color_buffer_half_float", "GL_EXT_copy_image", "GL_EXT_debug_label", "GL_EXT_debug_marker", "GL_EXT_discard_framebuffer", "GL_EXT_disjoint_timer_query", "GL_EXT_draw_buffers_indexed", "GL_EXT_external_buffer", "GL_EXT_geometry_shader", "GL_EXT_gpu_shader5", "GL_EXT_memory_object", "GL_EXT_memory_object_fd", "GL_EXT_multisampled_render_to_texture", "GL_EXT_multisampled_render_to_texture2", "GL_EXT_primitive_bounding_box", "GL_EXT_protected_textures", "GL_EXT_robustness", "GL_EXT_sRGB", "GL_EXT_sRGB_write_control", "GL_EXT_shader_framebuffer_fetch", "GL_EXT_shader_io_blocks", "GL_EXT_shader_non_constant_global_initializers", "GL_EXT_tessellation_shader", "GL_EXT_texture_border_clamp", "GL_EXT_texture_buffer", "GL_EXT_texture_cube_map_array", "GL_EXT_texture_filter_anisotropic", "GL_EXT_texture_format_BGRA8888", "GL_EXT_texture_format_sRGB_override", "GL_EXT_texture_norm16", "GL_EXT_texture_sRGB_R8", "GL_EXT_texture_sRGB_decode", "GL_EXT_texture_type_2_10_10_10_REV", "GL_KHR_blend_equation_advanced", "GL_KHR_blend_equation_advanced_coherent", "GL_KHR_debug", "GL_KHR_no_error", "GL_KHR_robust_buffer_access_behavior", "GL_KHR_texture_compression_astc_hdr", "GL_KHR_texture_compression_astc_ldr", "GL_NV_shader_noperspective_interpolation", "GL_OES_EGL_image", "GL_OES_EGL_image_external", "GL_OES_EGL_image_external_essl3", "GL_OES_EGL_sync", "GL_OES_blend_equation_separate", "GL_OES_blend_func_separate", "GL_OES_blend_subtract", "GL_OES_compressed_ETC1_RGB8_texture", "GL_OES_compressed_paletted_texture", "GL_OES_depth24", "GL_OES_depth_texture", "GL_OES_depth_texture_cube_map", "GL_OES_draw_texture", "GL_OES_element_index_uint", "GL_OES_framebuffer_object", "GL_OES_get_program_binary", "GL_OES_matrix_palette", "GL_OES_packed_depth_stencil", "GL_OES_point_size_array", "GL_OES_point_sprite", "GL_OES_read_format", "GL_OES_rgb8_rgba8", "GL_OES_sample_shading", "GL_OES_sample_variables", "GL_OES_shader_image_atomic", "GL_OES_shader_multisample_interpolation", "GL_OES_standard_derivatives", "GL_OES_stencil_wrap", "GL_OES_surfaceless_context", "GL_OES_texture_3D", "GL_OES_texture_compression_astc", "GL_OES_texture_cube_map", "GL_OES_texture_env_crossbar", "GL_OES_texture_float", "GL_OES_texture_float_linear", "GL_OES_texture_half_float", "GL_OES_texture_half_float_linear", "GL_OES_texture_mirrored_repeat", "GL_OES_texture_npot", "GL_OES_texture_stencil8", "GL_OES_texture_storage_multisample_2d_array", "GL_OES_vertex_array_object", "GL_OES_vertex_half_float", "GL_OVR_multiview", "GL_OVR_multiview2", "GL_OVR_multiview_multisampled_render_to_texture", "GL_QCOM_alpha_test", "GL_QCOM_extended_get", "GL_QCOM_shader_framebuffer_fetch_noncoherent", "GL_QCOM_texture_foveated", "GL_QCOM_tiled_rendering"},
protobuf.Tag{Number:16, Name:"varint"}:uint64(0),
protobuf.Tag{Number:9, Name:"message"}:protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:uint64(8371739161332442222)},
protobuf.Tag{Number:11, Name:"string"}:[]string{"arm64-v8a", "armeabi-v7a", "armeabi"},
protobuf.Tag{Number:13, Name:"varint"}:uint64(2073),
protobuf.Tag{Number:21, Name:"varint"}:uint64(8),
protobuf.Tag{Number:2, Name:"varint"}:uint64(1),
protobuf.Tag{Number:3, Name:"varint"}:uint64(1),
protobuf.Tag{Number:10, Name:"message"}:[]protobuf.Message{protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:[]uint64{0x682e64696f72646e, 0x632e657261776472, 0x796e612e6172656d}}, protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:[]uint64{0x682e64696f72646e, 0x6d2e657261776472},
protobuf.Tag{Number:13, Name:"fixed64"}:uint64(7308901739622527587)}, protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:[]uint64{0x682e64696f72646e, 0x722e657261776472, 0x6c616d726f6e2e6d}}, protobuf.Message{protobuf.Tag{Number:12, Name:"fixed64"}:[]uint64{0x682e64696f72646e, 0x772e657261776472},
protobuf.Tag{Number:13, Name:"fixed64"}:uint64(7310012310535170406)}},
protobuf.Tag{Number:14, Name:"message"}:[]protobuf.Message{protobuf.Message{}, protobuf.Message{}, protobuf.Message{}, protobuf.Message{}, protobuf.Message{}, protobuf.Message{}, protobuf.Message{}, protobuf.Message{}, protobuf.Message{}, protobuf.Message{}, protobuf.Message{protobuf.Tag{Number:13, Name:"varint"}:uint64(105)}, protobuf.Message{protobuf.Tag{Number:13, Name:"varint"}:uint64(114)}, protobuf.Message{protobuf.Tag{Number:13, Name:"varint"}:uint64(117)}, protobuf.Message{protobuf.Tag{Number:13, Name:"varint"}:uint64(121)}, protobuf.Message{}, protobuf.Message{protobuf.Tag{Number:14, Name:"varint"}:uint64(97)}, protobuf.Message{protobuf.Tag{Number:14, Name:"varint"}:uint64(108)}, protobuf.Message{protobuf.Tag{Number:14, Name:"varint"}:uint64(116)}, protobuf.Message{}},
protobuf.Tag{Number:19, Name:"varint"}:uint64(0),
protobuf.Tag{Number:20, Name:"varint"}:uint64(8589935000)},
protobuf.Tag{Number:20, Name:"varint"}:uint64(0),
protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:4, Name:"message"}:protobuf.Message{protobuf.Tag{Number:9, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"message"}:protobuf.Message{protobuf.Tag{Number:2, Name:"string"}:"sargo",
protobuf.Tag{Number:3, Name:"string"}:"google",
protobuf.Tag{Number:7, Name:"varint"}:uint64(1644766323),
protobuf.Tag{Number:14, Name:"varint"}:uint64(0),
protobuf.Tag{Number:1, Name:"string"}:"PQ3B.190705.003",
protobuf.Tag{Number:8, Name:"varint"}:uint64(203615028),
protobuf.Tag{Number:10, Name:"varint"}:uint64(29),
protobuf.Tag{Number:6, Name:"string"}:"android-google",
protobuf.Tag{Number:13, Name:"string"}:"sargo",
protobuf.Tag{Number:5, Name:"string"}:"b4s4-0.1-5613380",
protobuf.Tag{Number:9, Name:"string"}:"sargo",
protobuf.Tag{Number:11, Name:"string"}:"Pixel 3a",
protobuf.Tag{Number:12, Name:"string"}:"google",
protobuf.Tag{Number:4, Name:"string"}:"g670-00011-190411-B-5457439"},
protobuf.Tag{Number:2, Name:"varint"}:uint64(0),
protobuf.Tag{Number:6, Name:"string"}:"334050",
protobuf.Tag{Number:7, Name:"string"}:"20815",
protobuf.Tag{Number:8, Name:"string"}:"mobile-notroaming"},
protobuf.Tag{Number:6, Name:"message"}:protobuf.Message{},
protobuf.Tag{Number:6, Name:"string"}:"en_GB"}

func main() {
   var req http.Request
   req.Body = io.NopCloser(bytes.NewReader(body.Marshal()))
   req.Header = make(http.Header)
   req.Header["App"] = []string{"com.google.android.gms"}
   req.Header["Content-Type"] = []string{"application/x-protobuffer"}
   req.Header["Host"] = []string{"android.clients.google.com"}
   req.Header["User-Agent"] = []string{"GoogleAuth/1.4 sargo PQ3B.190705.003"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/checkin"
   req.URL.Scheme = "http"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   mes, err := protobuf.Decode(res.Body)
   if err != nil {
      panic(err)
   }
   androidID := mes.GetFixed64(7, "androidId")
   googleplay.LogLevel = 1
   tok := googleplay.Token{token}
   auth, err := tok.Auth()
   if err != nil {
      panic(err)
   }
   dev := googleplay.Device{androidID}
   det, err := auth.Header(&dev, false).Details("com.bbca.bbcafullepisodes")
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", det)
}
