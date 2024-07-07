import { Component } from '@angular/core';
import { NzListModule } from 'ng-zorro-antd/list';
import { result } from '../../class/result';
import { ResultService } from '../services/result.service';
import { response } from '../../class/response';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import { NzButtonModule } from 'ng-zorro-antd/button';

@Component({
  selector: 'app-result',
  standalone: true,
  imports: [NzListModule,NzButtonModule],
  templateUrl: './result.component.html',
  styleUrl: './result.component.css'
})
export class ResultComponent {
  results:result[]=[];
  username:string="ysm"   // 等实现登陆后获取
  loadingMore = false;
  page:number=1;

  constructor(
    private resultService:ResultService,
    private notification:NzNotificationService,
  ){}

  ngOnInit(): void {
    this.getAllResultByUser(this.page)
  }

  getAllResultByUser(page:number){
    this.resultService.getResultByU(this.username,page).subscribe((r:response)=>{
      if(r.status!==0){
        this.notification.error("获取结果失败",r.msg);
        return;
      }
      if (r.data.length=== 0){
        this.loadingMore = false;
        return;
      }
      this.results = [...this.results,...r.data]
    })
  }

  onLoadMore(){
    this.page = this.page+1;
    this.getAllResultByUser(this.page)
  }
}