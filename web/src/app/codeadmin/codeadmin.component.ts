import { Component } from '@angular/core';
import { NzTableModule } from 'ng-zorro-antd/table';
import { Person, script } from '../../class/script';
import { CommonModule } from '@angular/common';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { RouterModule } from '@angular/router';
import { ScriptService } from '../services/script.service';
import { response } from '../../class/response';
import { NzNotificationModule, NzNotificationService } from 'ng-zorro-antd/notification';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';



@Component({
  selector: 'app-codeadmin',
  standalone: true,
  imports: [
    NzTableModule,
    CommonModule,
    NzButtonModule,
    RouterModule,
    NzNotificationModule,
    NzPopconfirmModule
  ],
  templateUrl: './codeadmin.component.html',
  styleUrl: './codeadmin.component.css'
})

export class CodeadminComponent {
  listOfData: Person[] = [];

  scripts:script[]=[];
  private username:string="ysm"

  constructor(
    private scriptService: ScriptService,
    private notification:NzNotificationService,
  ){}

  ngOnInit(): void {
    const s:script={
      username:this.username
    }
    this.scriptService.getAllScript(s).subscribe((r:response)=>{
      if (r.status!==0){
        this.notification.error("系统bug",r.msg);
        return;
      }
      this.scripts = r.data
    })
  }

  delScript(id:any){
    const s:script={id:id}
    this.scriptService.delScript(s).subscribe((r:response)=>{
      if (r.status!==0){
        this.notification.error("系统bug",r.msg);
        return;
      }
      this.notification.error("恭喜","删除成功");
    })
  }
}
