<template>
  <div>
    <el-form :inline="true">
      <el-form-item>
        <el-button type="primary" class="copy" icon="el-icon-document-copy" size="small" :data-clipboard-text="ssl.key" @click="copy">复制私钥</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" class="copy" icon="el-icon-document-copy" size="small" :data-clipboard-text="ssl.pem" @click="copy">复制公钥</el-button>
      </el-form-item>
      <el-form-item>
        <el-button icon="el-icon-download" size="small" @click="download('key')">下载私钥</el-button>
      </el-form-item>
      <el-form-item>
        <el-button icon="el-icon-download" size="small" @click="download('pem')">下载公钥</el-button>
      </el-form-item>
      <el-form-item>
        <span style="color: #3d763e;">证书到期时间: 2022-10-28 07:59:59</span>
      </el-form-item>
    </el-form>
    <prism-editor class="my-editor" :value="ssl.key" readonly :highlight="highlighter" line-numbers style="width: 48%;float: left;"></prism-editor>
    <prism-editor class="my-editor" :value="ssl.pem" readonly :highlight="highlighter" line-numbers style="width: 48%;float: right;"></prism-editor>
  </div>
</template>

<script>
import {PrismEditor} from 'vue-prism-editor';
import 'vue-prism-editor/dist/prismeditor.min.css';
import {highlight, languages} from 'prismjs/components/prism-core';
import 'prismjs/components/prism-clike';
import 'prismjs/components/prism-javascript';
import 'prismjs/themes/prism-tomorrow.css';
import Clipboard from "clipboard";
import {Message} from "element-ui";

export default {
  components: {
    PrismEditor,
  },
  data() {
    return {
      ssl:{
        key: '-----BEGIN RSA PRIVATE KEY-----\n' +
            'MIIEpQIBAAKCAQEAnvZ19F6feGHWD17ID9ChOGuqwpuTMNMD2cwK0Q6Xe/qmxvfu\n' +
            'ODHi4rMwTwT6wIP641D8+EWjs8xUvmsnR//0ewL0xkVvJvP95jtov/A1c7tuDQft\n' +
            'rEFVzCHxc/kvExc+ta7jUteP3Ul/kLpp9uxOf1upCkY71M1UhlK5RjN+ZPXtEMg1\n' +
            'B4YxC9NEzL3mt266hdV1ytm6NY11QpdfAG74ROqUHC7M2iRlSaBRZ+QuaebAz+OC\n' +
            'VVCsJRh1hQgJV482ndIFGAvGXxoigOvGCky68A2nvyHGPrJn0vguw2/iPL4QmQdk\n' +
            'rrOTSYo2nDsy1KQH1fKzwDGwxZAl/s32/V1bSwIDAQABAoIBAEnNTWg+QjWm3rZW\n' +
            'H/ZPu3oxap9dkZHadnA47lCvQJ97+NBprX0DlD4CSYx2WGDnJaVCJaCy3FUOqc2a\n' +
            'fCtOZ1hxP+mJtRhVsWyh6nkqcfagb3C3XJ5b2xSiMbLRXwyy/xdk8f+1OwLXK9Wq\n' +
            '5qWQgsVYbfvqLb8X8YVHhONV5vnCBI9xjcFzQll9hGXhS/VyAEZONm0/9V8efeAL\n' +
            'f/eYxcIVcVy0swseaU0XsX2d6WzdXNya3PJjpxBDB3qmtAUTTHeGEAENqSRSE0Wb\n' +
            'cvecnyBC9I8IROpvciuO703UP8Ufc4DXauTRRbVcXr+X5lkU3SVPKWpOqphuk4A+\n' +
            '69oPeDkCgYEAzKba2kq3ydETHhIAF5bHqbSDv3qnAYvTYaeV7IiBpvJK90IslmDn\n' +
            'oR4HEIbbx0r19YdpKv9hhnUmZLb0PxeA4Kxne7d59TJOlNJamQ6UpyuoQjL1yhUl\n' +
            'w1B8VC2K6+1pW6Kl46OuxzzJpsqY7N3dHnNxwx49LOIbVda4hqNbmUcCgYEAxtjr\n' +
            'DCVv+rRv9DgqTbZHsROyv4wokb/hhyUJIxXk6NXw4laN6fqVe+S22x06hawlwoZc\n' +
            'sgpVRp5nCI3QFrb2Ytew/vS+dTkeg8gjaEX0nXkk8z2h6ykzjfYV5A2DljTz0+E/\n' +
            'IwcYu5q2wCnjCqGuJgltL+yOwx2gWHVEMTuqL90CgYEAq/twoc/2wahs+3E/GJPC\n' +
            'uRAVQ73sSTVbb46pRHy3EDg/3aiD1eAb599XKoFagxGwvySLsfGp4JfkqvDBM6Ap\n' +
            '0yWvVV2sJO2g8hoQvt1s+UuL/A3dVnUNJiUmRot7WqdKShLWfAYbJB4mjH9nom8U\n' +
            'mrmDmeh4mpZtSGbI7YSsvpMCgYEAwVNCasO3NFLbieBVWMrAtuvE3LBC/TlF5mPe\n' +
            'l0sBDFJ8aUHnpbVoQlUiZrPrIu00pQSG+N5M/GSjHdDqoDrwHzc/MRH4XfXTsiIS\n' +
            'Mp3Y7oVKpGamlwCuy/WWYOu2pIIFEVmsTbznkfYgeYH/6GU68GY3zuE6B381Vutk\n' +
            'Ck3M780CgYEAtR3GBuSjLJl+fXBFnUzgtB2Imx64TViG1GDVKV3GHD7njfUn3PJw\n' +
            '23M1MnrPpOrvh2IXmiDf81wfGLOT7ipk1l7PffuhPiP/yB2C1AgX0nr3QEDmukY/\n' +
            'VUbXkLKrkEZoQ29rRnzdNs40fGhKanaMoPoVQ03LMFINrJqH0aZdRWk=\n' +
            '-----END RSA PRIVATE KEY-----\n',
        pem: '-----BEGIN CERTIFICATE-----\n' +
            'MIIGfjCCBGagAwIBAgIQPlGxglvD5649j9S8so7PsTANBgkqhkiG9w0BAQwFADBL\n' +
            'MQswCQYDVQQGEwJBVDEQMA4GA1UEChMHWmVyb1NTTDEqMCgGA1UEAxMhWmVyb1NT\n' +
            'TCBSU0EgRG9tYWluIFNlY3VyZSBTaXRlIENBMB4XDTIyMDcyOTAwMDAwMFoXDTIy\n' +
            'MTAyNzIzNTk1OVowFDESMBAGA1UEAwwJKi5uaWttLmNuMIIBIjANBgkqhkiG9w0B\n' +
            'AQEFAAOCAQ8AMIIBCgKCAQEAnvZ19F6feGHWD17ID9ChOGuqwpuTMNMD2cwK0Q6X\n' +
            'e/qmxvfuODHi4rMwTwT6wIP641D8+EWjs8xUvmsnR//0ewL0xkVvJvP95jtov/A1\n' +
            'c7tuDQftrEFVzCHxc/kvExc+ta7jUteP3Ul/kLpp9uxOf1upCkY71M1UhlK5RjN+\n' +
            'ZPXtEMg1B4YxC9NEzL3mt266hdV1ytm6NY11QpdfAG74ROqUHC7M2iRlSaBRZ+Qu\n' +
            'aebAz+OCVVCsJRh1hQgJV482ndIFGAvGXxoigOvGCky68A2nvyHGPrJn0vguw2/i\n' +
            'PL4QmQdkrrOTSYo2nDsy1KQH1fKzwDGwxZAl/s32/V1bSwIDAQABo4ICkzCCAo8w\n' +
            'HwYDVR0jBBgwFoAUyNl4aKLZGWjVPXLeXwo+3LWGhqYwHQYDVR0OBBYEFFoN+hqR\n' +
            'tJDFkypFFneA98jMsMcmMA4GA1UdDwEB/wQEAwIFoDAMBgNVHRMBAf8EAjAAMB0G\n' +
            'A1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjBJBgNVHSAEQjBAMDQGCysGAQQB\n' +
            'sjEBAgJOMCUwIwYIKwYBBQUHAgEWF2h0dHBzOi8vc2VjdGlnby5jb20vQ1BTMAgG\n' +
            'BmeBDAECATCBiAYIKwYBBQUHAQEEfDB6MEsGCCsGAQUFBzAChj9odHRwOi8vemVy\n' +
            'b3NzbC5jcnQuc2VjdGlnby5jb20vWmVyb1NTTFJTQURvbWFpblNlY3VyZVNpdGVD\n' +
            'QS5jcnQwKwYIKwYBBQUHMAGGH2h0dHA6Ly96ZXJvc3NsLm9jc3Auc2VjdGlnby5j\n' +
            'b20wggEFBgorBgEEAdZ5AgQCBIH2BIHzAPEAdwBGpVXrdfqRIDC1oolp9PN9ESxB\n' +
            'dL79SbiFq/L8cP5tRwAAAYJKzkI2AAAEAwBIMEYCIQDaJ/WJq8HOMMJX5zJJ5kwn\n' +
            'buMRTA5rLiLEhB63Y4mMzwIhAILRBM1KHy6IozlOwzHcPxX50r7MneS7iMoN/pcY\n' +
            '1UmlAHYAQcjKsd8iRkoQxqE6CUKHXk4xixsD6+tLx2jwkGKWBvYAAAGCSs5CSgAA\n' +
            'BAMARzBFAiEA8FDwSckYYfpi0pDFGi+x0US49KS+E4GriotdHs6qPYwCIG1C+2lU\n' +
            'dvw4gmrhFa8YgzmIQRa05Ks00HGF/D5erwAoMDEGA1UdEQQqMCiCCSoubmlrbS5j\n' +
            'boIJKi5odHRvLmNuggdodHRvLmNuggduaWttLmNuMA0GCSqGSIb3DQEBDAUAA4IC\n' +
            'AQBLPR78F8Ieo9QI+ugZMDU0CuwyCs4gcwXK3N1CTqz26sdXxbUygXs5SQfhNTuc\n' +
            '9jua21o7PSIia+Af+kWUNLRK2khYxipd9gv4TXjHpGZeYVq5O70pZlwXKYYQ9w6O\n' +
            'EJVA8WIUNi9Oo+QVsPt9BSzrsOOiwaYFTe/WfDZaBhtwgeye/agf7gcdfqMVkCns\n' +
            'tjqbP8n7kxuO13Xhun9a/ze563uXFkzm3i3kF/Nzg1RuARKZFAeVNkEnJXT6IcbN\n' +
            'O5R5Kf7fzbIJL7mZatF9uwr/n74lKnvJS9F5UgX1PrhgTRb5D8Naagaq5b2R0NZU\n' +
            '3ySeabOJ4l5NSMG2Qaex2EaqwVTfyDKOkF6PxD7kdP38SlZZnbsqp/i2o/A++En4\n' +
            'bDAZAhBNIYtGWMJhxikgQyQe1V+ZJJCpvITUZLkKV0n2kATQZjbHXx7EJ8PqZ5CG\n' +
            'Y7Ye0oxGHUv6bTndh3qzRmz+Vu4nmOI+HMA1NrFv8j3oiX8iRSSAiyZQnrqzka+y\n' +
            'zi+UdniAeR1mu/KmcbFYN8hA94zX0VOBXGipB5pbJTgPZzt9WMtzldQKQReV15xF\n' +
            '6ZsIxkSQ3SoW/HtNndj2H0SBHiLvVhKrF6VAUAubDDfYZLp0jmsXZ/z8kxC9ozMg\n' +
            'mvRoInfrxwTToJDyFgf7jhbEAESbT/Qv7K1TR3yGjT28zg==\n' +
            '-----END CERTIFICATE-----\n' +
            '-----BEGIN CERTIFICATE-----\n' +
            'MIIG1TCCBL2gAwIBAgIQbFWr29AHksedBwzYEZ7WvzANBgkqhkiG9w0BAQwFADCB\n' +
            'iDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCk5ldyBKZXJzZXkxFDASBgNVBAcTC0pl\n' +
            'cnNleSBDaXR5MR4wHAYDVQQKExVUaGUgVVNFUlRSVVNUIE5ldHdvcmsxLjAsBgNV\n' +
            'BAMTJVVTRVJUcnVzdCBSU0EgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMjAw\n' +
            'MTMwMDAwMDAwWhcNMzAwMTI5MjM1OTU5WjBLMQswCQYDVQQGEwJBVDEQMA4GA1UE\n' +
            'ChMHWmVyb1NTTDEqMCgGA1UEAxMhWmVyb1NTTCBSU0EgRG9tYWluIFNlY3VyZSBT\n' +
            'aXRlIENBMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAhmlzfqO1Mdgj\n' +
            '4W3dpBPTVBX1AuvcAyG1fl0dUnw/MeueCWzRWTheZ35LVo91kLI3DDVaZKW+TBAs\n' +
            'JBjEbYmMwcWSTWYCg5334SF0+ctDAsFxsX+rTDh9kSrG/4mp6OShubLaEIUJiZo4\n' +
            't873TuSd0Wj5DWt3DtpAG8T35l/v+xrN8ub8PSSoX5Vkgw+jWf4KQtNvUFLDq8mF\n' +
            'WhUnPL6jHAADXpvs4lTNYwOtx9yQtbpxwSt7QJY1+ICrmRJB6BuKRt/jfDJF9Jsc\n' +
            'RQVlHIxQdKAJl7oaVnXgDkqtk2qddd3kCDXd74gv813G91z7CjsGyJ93oJIlNS3U\n' +
            'gFbD6V54JMgZ3rSmotYbz98oZxX7MKbtCm1aJ/q+hTv2YK1yMxrnfcieKmOYBbFD\n' +
            'hnW5O6RMA703dBK92j6XRN2EttLkQuujZgy+jXRKtaWMIlkNkWJmOiHmErQngHvt\n' +
            'iNkIcjJumq1ddFX4iaTI40a6zgvIBtxFeDs2RfcaH73er7ctNUUqgQT5rFgJhMmF\n' +
            'x76rQgB5OZUkodb5k2ex7P+Gu4J86bS15094UuYcV09hVeknmTh5Ex9CBKipLS2W\n' +
            '2wKBakf+aVYnNCU6S0nASqt2xrZpGC1v7v6DhuepyyJtn3qSV2PoBiU5Sql+aARp\n' +
            'wUibQMGm44gjyNDqDlVp+ShLQlUH9x8CAwEAAaOCAXUwggFxMB8GA1UdIwQYMBaA\n' +
            'FFN5v1qqK0rPVIDh2JvAnfKyA2bLMB0GA1UdDgQWBBTI2XhootkZaNU9ct5fCj7c\n' +
            'tYaGpjAOBgNVHQ8BAf8EBAMCAYYwEgYDVR0TAQH/BAgwBgEB/wIBADAdBgNVHSUE\n' +
            'FjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwIgYDVR0gBBswGTANBgsrBgEEAbIxAQIC\n' +
            'TjAIBgZngQwBAgEwUAYDVR0fBEkwRzBFoEOgQYY/aHR0cDovL2NybC51c2VydHJ1\n' +
            'c3QuY29tL1VTRVJUcnVzdFJTQUNlcnRpZmljYXRpb25BdXRob3JpdHkuY3JsMHYG\n' +
            'CCsGAQUFBwEBBGowaDA/BggrBgEFBQcwAoYzaHR0cDovL2NydC51c2VydHJ1c3Qu\n' +
            'Y29tL1VTRVJUcnVzdFJTQUFkZFRydXN0Q0EuY3J0MCUGCCsGAQUFBzABhhlodHRw\n' +
            'Oi8vb2NzcC51c2VydHJ1c3QuY29tMA0GCSqGSIb3DQEBDAUAA4ICAQAVDwoIzQDV\n' +
            'ercT0eYqZjBNJ8VNWwVFlQOtZERqn5iWnEVaLZZdzxlbvz2Fx0ExUNuUEgYkIVM4\n' +
            'YocKkCQ7hO5noicoq/DrEYH5IuNcuW1I8JJZ9DLuB1fYvIHlZ2JG46iNbVKA3ygA\n' +
            'Ez86RvDQlt2C494qqPVItRjrz9YlJEGT0DrttyApq0YLFDzf+Z1pkMhh7c+7fXeJ\n' +
            'qmIhfJpduKc8HEQkYQQShen426S3H0JrIAbKcBCiyYFuOhfyvuwVCFDfFvrjADjd\n' +
            '4jX1uQXd161IyFRbm89s2Oj5oU1wDYz5sx+hoCuh6lSs+/uPuWomIq3y1GDFNafW\n' +
            '+LsHBU16lQo5Q2yh25laQsKRgyPmMpHJ98edm6y2sHUabASmRHxvGiuwwE25aDU0\n' +
            '2SAeepyImJ2CzB80YG7WxlynHqNhpE7xfC7PzQlLgmfEHdU+tHFeQazRQnrFkW2W\n' +
            'kqRGIq7cKRnyypvjPMkjeiV9lRdAM9fSJvsB3svUuu1coIG1xxI1yegoGM4r5QP4\n' +
            'RGIVvYaiI76C0djoSbQ/dkIUUXQuB8AL5jyH34g3BZaaXyvpmnV4ilppMXVAnAYG\n' +
            'ON51WhJ6W0xNdNJwzYASZYH+tmCWI+N60Gv2NNMGHwMZ7e9bXgzUCZH5FaBFDGR5\n' +
            'S9VWqHB73Q+OyIVvIbKYcSc2w/aSuFKGSA==\n' +
            '-----END CERTIFICATE-----\n' +
            '-----BEGIN CERTIFICATE-----\n' +
            'MIIFgTCCBGmgAwIBAgIQOXJEOvkit1HX02wQ3TE1lTANBgkqhkiG9w0BAQwFADB7\n' +
            'MQswCQYDVQQGEwJHQjEbMBkGA1UECAwSR3JlYXRlciBNYW5jaGVzdGVyMRAwDgYD\n' +
            'VQQHDAdTYWxmb3JkMRowGAYDVQQKDBFDb21vZG8gQ0EgTGltaXRlZDEhMB8GA1UE\n' +
            'AwwYQUFBIENlcnRpZmljYXRlIFNlcnZpY2VzMB4XDTE5MDMxMjAwMDAwMFoXDTI4\n' +
            'MTIzMTIzNTk1OVowgYgxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpOZXcgSmVyc2V5\n' +
            'MRQwEgYDVQQHEwtKZXJzZXkgQ2l0eTEeMBwGA1UEChMVVGhlIFVTRVJUUlVTVCBO\n' +
            'ZXR3b3JrMS4wLAYDVQQDEyVVU0VSVHJ1c3QgUlNBIENlcnRpZmljYXRpb24gQXV0\n' +
            'aG9yaXR5MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAgBJlFzYOw9sI\n' +
            's9CsVw127c0n00ytUINh4qogTQktZAnczomfzD2p7PbPwdzx07HWezcoEStH2jnG\n' +
            'vDoZtF+mvX2do2NCtnbyqTsrkfjib9DsFiCQCT7i6HTJGLSR1GJk23+jBvGIGGqQ\n' +
            'Ijy8/hPwhxR79uQfjtTkUcYRZ0YIUcuGFFQ/vDP+fmyc/xadGL1RjjWmp2bIcmfb\n' +
            'IWax1Jt4A8BQOujM8Ny8nkz+rwWWNR9XWrf/zvk9tyy29lTdyOcSOk2uTIq3XJq0\n' +
            'tyA9yn8iNK5+O2hmAUTnAU5GU5szYPeUvlM3kHND8zLDU+/bqv50TmnHa4xgk97E\n' +
            'xwzf4TKuzJM7UXiVZ4vuPVb+DNBpDxsP8yUmazNt925H+nND5X4OpWaxKXwyhGNV\n' +
            'icQNwZNUMBkTrNN9N6frXTpsNVzbQdcS2qlJC9/YgIoJk2KOtWbPJYjNhLixP6Q5\n' +
            'D9kCnusSTJV882sFqV4Wg8y4Z+LoE53MW4LTTLPtW//e5XOsIzstAL81VXQJSdhJ\n' +
            'WBp/kjbmUZIO8yZ9HE0XvMnsQybQv0FfQKlERPSZ51eHnlAfV1SoPv10Yy+xUGUJ\n' +
            '5lhCLkMaTLTwJUdZ+gQek9QmRkpQgbLevni3/GcV4clXhB4PY9bpYrrWX1Uu6lzG\n' +
            'KAgEJTm4Diup8kyXHAc/DVL17e8vgg8CAwEAAaOB8jCB7zAfBgNVHSMEGDAWgBSg\n' +
            'EQojPpbxB+zirynvgqV/0DCktDAdBgNVHQ4EFgQUU3m/WqorSs9UgOHYm8Cd8rID\n' +
            'ZsswDgYDVR0PAQH/BAQDAgGGMA8GA1UdEwEB/wQFMAMBAf8wEQYDVR0gBAowCDAG\n' +
            'BgRVHSAAMEMGA1UdHwQ8MDowOKA2oDSGMmh0dHA6Ly9jcmwuY29tb2RvY2EuY29t\n' +
            'L0FBQUNlcnRpZmljYXRlU2VydmljZXMuY3JsMDQGCCsGAQUFBwEBBCgwJjAkBggr\n' +
            'BgEFBQcwAYYYaHR0cDovL29jc3AuY29tb2RvY2EuY29tMA0GCSqGSIb3DQEBDAUA\n' +
            'A4IBAQAYh1HcdCE9nIrgJ7cz0C7M7PDmy14R3iJvm3WOnnL+5Nb+qh+cli3vA0p+\n' +
            'rvSNb3I8QzvAP+u431yqqcau8vzY7qN7Q/aGNnwU4M309z/+3ri0ivCRlv79Q2R+\n' +
            '/czSAaF9ffgZGclCKxO/WIu6pKJmBHaIkU4MiRTOok3JMrO66BQavHHxW/BBC5gA\n' +
            'CiIDEOUMsfnNkjcZ7Tvx5Dq2+UUTJnWvu6rvP3t3O9LEApE9GQDTF1w52z97GA1F\n' +
            'zZOFli9d31kWTz9RvdVFGD/tSo7oBmF0Ixa1DVBzJ0RHfxBdiSprhTEUxOipakyA\n' +
            'vGp4z7h/jnZymQyd/teRCBaho1+V\n' +
            '-----END CERTIFICATE-----\n',
      },
    }
  },
  methods: {
    download(type){
      let export_blob = new Blob([this.ssl[type]])
      let save_link = document.createElement("a")
      save_link.href = window.URL.createObjectURL(export_blob)
      save_link.download = type +'.txt'
      save_link.click()
    },
    copy(){
      let clipboard = new Clipboard(".copy")
      clipboard.on('success', e => {
        Message({
          type: 'success',
          message: '复制成功',
          showClose: true
        })
        clipboard.destroy()
      })
      clipboard.on('error', e => {
        Message({
          type: 'error',
          message: '复制失败,',
          showClose: true,
        })
        clipboard.destroy()
      })
    },
    highlighter(code) {
      return highlight(code, languages.plaintext)
    },
  },
}
</script>

<style scoped>
.my-editor {
  background: #fafafa;
  font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
  font-size: 14px;
  line-height: 1.5;
  padding: 5px 5px 5px 10px;
  border: 1px solid #e2e2e2;
}
</style>
<style>
.my-editor textarea:focus {
  outline: none;
}
</style>
