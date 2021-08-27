# 台灣適用宗教系統

## 主要目前是台灣佛教相關法會系統，無金錢相關記錄
整個系統單機執行，無上網需求，所有前後端都是可以單獨執行，不需要依賴線上檔案，避免將來產生一堆問題，讓系統可以永遠執行

因為前後端分離情況下，前端可以在不影響後端的情況，前端可以做很多調整

## 運行方式
1. 自行 Compiler 產生 main.exe，產生後先執行 main.exe  
2. 執行瀏灠器輸入 http://localhost:8081/ 

## 系統使用
1. Vue.js v2.6.14
2. Bootstrap v4.1.1
3. Bootstrap-vue
4. golang
5. sqlite

## Golang Sqlite Compiler
### Windows
https://hoohoo.top/blog/golang-fix-gccexec-gcc-executable-file-not-found-in-path/  
在 Windows 如果使用 Go 語言使用 sqlite3 時，會需要透過 gcd 來編譯 C  ... 請看連結  
新版 tdm64-gcc-10.3.0-2.exe  
https://jmeubank.github.io/tdm-gcc/articles/2021-05/10.3.0-release

## globalVarialbe.js
href 必需設定正確

## 導灠
[導灠1](導灠1.mp4)

## 會員編號
建議 年月日4碼流水編：202107270001、202107270002、202107270003 ...  

## DB 結構

會員    member  
點燈    diandeng  (目前未使用)  
長祈    zhangqi  (長生祈福 <=> 消災祈福)  
超薦個人  chaojiangeren  
超薦祖先  chaojianzuxian  
嬰靈    yingling  
地基主  dijizhu  
動物靈  dongwuling  


### 會員 member  
m_num  
m_name  
m_tel  
m_addr  
m_diandeng  
m_zhangqi  
m_chaojiangeren  
m_chaojianzuxian  
m_yingling  
m_dongwuling  
m_dijizhu  
