import web_webview from '@ohos.web.webview'
import router from '@ohos.router'

@Entry
@Component
struct web{

  // 获取页面参数
  @State params:any = router.getParams()

  controller: web_webview.WebviewController = new web_webview.WebviewController()
  @State mode: MixedMode = MixedMode.All
  build() {
    Column() {
      Row(){
          Text("返回").fontWeight(FontWeight.Bold)
            .onClick(()=>{
              // 返回上一页
              router.back()
            })
      }
      .width("100%")
      .height(30)
      .margin(10)
      .padding(5)

      if(this.params){
        Text(this.params.header).fontWeight(FontWeight.Bold).width("100%")
      }

      Web({ src: '', controller: this.controller })
        .mixedMode(this.mode)
    }
    .width("100%")
    .height("100%")
  }



}