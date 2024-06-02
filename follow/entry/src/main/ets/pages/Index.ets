import router from '@ohos.router'
import result from '../../viewmodel/result'
import http from '@ohos.net.http';
import ResultModel from '../../model/ResultModel'


// 后端传递参数格式
class item {
  header:string   // 标题
  image:Resource    // 图像
  content:string  // 内容

  constructor(header:string,image:Resource,content:string) {
    this.header = header
    this.image = image
    this.content = content
  }
}


@Entry
@Component
struct Index {

  // 构造简单数据
  // private items:Array<item> = [
  //   new item("动漫1",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫2",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫3",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫4",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫5",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫6",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫7",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫8",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫9",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫10",$r("app.media.mei"),"这个动漫女生美吗？"),
  //   new item("动漫11",$r("app.media.mei"),"这个动漫女生美吗？")
  // ]

  @State results: result[] = []
  aboutToAppear(){
    this.loadResults()
  }

  @Builder
  SubMenu() {
    Menu() {
      MenuItem({ content: "复制", labelInfo: "Ctrl+C" })
      MenuItem({ content: "粘贴", labelInfo: "Ctrl+V" })
    }
  }

  @Builder
  ClassMenu(){
    Menu(){
      MenuItem({ content: "全部"})
      MenuItem({
        content: "动漫",
        builder: this.SubMenu.bind(this)
      })
      MenuItem({ content: "招聘"})
      MenuItem({ content: "电影"})
      MenuItem({ content: "音乐"})
    }
  }

  build() {
    Column(){
      Row(){
        Column() {
          Text('菜单')
            .fontWeight(FontWeight.Bold)
            .fontColor(Color.White)
        }
        .bindMenu(this.ClassMenu)
        .margin({left:20})
      }
      .backgroundColor(Color.Black)
      .width("100%")
      .height(60)


      List({space:10}){
        ForEach(this.results,(item:item)=>{
          ListItem(){
            Row(){
              Image(item.image).width("150")
              Column({space:4}){
                Text(item.header)
                  .fontSize(20)
                  .fontWeight(FontWeight.Bold)
                  .margin({bottom:30})
                Text(item.content)
                  .fontSize(16)
                  .fontWeight(FontWeight.Medium)
              }
              .alignItems(HorizontalAlign.Start)
              .margin({left:15})
              .height("100%")
            }
            .width("100%")
            .backgroundColor(0xFFFFFF)
            .borderRadius(20)
            .height(120)
            .padding(10)
            .onClick(()=>{
              router.pushUrl(
                {
                  url:"pages/web",
                  params:{
                    header:item.header
                  }
                },
                router.RouterMode.Single,
                err => {
                  if (err){
                    console.log("路由失败，errCode:",err.code,"msg:",err.message)
                  }
                }
              )
            })
          }
        })
      }
      .width("100%")
      .layoutWeight(1)
      .backgroundColor(Color.Black)
      .padding(10)
    }.width("100%")
  }


  loadResults(){
    this.getResults()
      .then(results => {
        this.results = results;
        console.log("this is a test",JSON.stringify(this.results));
      })
  }

  getResults():Promise<result[]>{
    return new Promise((resolve,reject)=>{
      let urlPath = "http://localhost:8080/getAllResult?username=follow&pwd=follow@123456"
      let httpRequest = http.createHttp()
      httpRequest.request(
        urlPath,
        {
          method:http.RequestMethod.GET,
          header:{
            'Content-Type': 'application/json',
          }
        },
      ).then(resp => {
        if(resp.responseCode === 200){
          console.log("查询成功");
          resolve(JSON.parse(resp.result.toString()))
        }else{
          console.log("查询失败，error:",JSON.stringify(resp));
          reject("查询失败")
        }
      })
        .catch(error => {
          console.log("查询爬虫结果失败！error:",JSON.stringify(error));
        })
    })
  }
}