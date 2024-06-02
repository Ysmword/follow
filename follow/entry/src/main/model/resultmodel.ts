
import http from '@ohos.net.http';
import result from '../viewmodel/result';


class ResultModel{

  getResults():Promise<result[]>{
    return new Promise((resolve,reject)=>{
      let urlPath = "http://localhost:8080/getAllResult?username=follow&pwd=follow@123456"
      let httpRequest = http.createHttp()
      httpRequest.request(
        urlPath,
        {
          method:http.RequestMethod.GET,
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

const resultModel = new ResultModel();

export default resultModel as ResultModel;
