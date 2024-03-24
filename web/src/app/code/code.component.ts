import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzFormModule } from 'ng-zorro-antd/form';
import { CodemirrorModule } from '@ctrl/ngx-codemirror';
import { FormControl, FormGroup, FormsModule, NonNullableFormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzGridModule } from 'ng-zorro-antd/grid';
import 'codemirror/theme/blackboard.css';
import { ScriptService } from '../services/script.service';
import { script } from '../../class/script';
import { response } from '../../class/response';
import { NzNotificationModule, NzNotificationService } from 'ng-zorro-antd/notification';
import { NzModalModule, NzModalService } from 'ng-zorro-antd/modal';


@Component({
  selector: 'app-code',
  standalone: true,
  imports: [
    NzInputModule,
    NzSelectModule,
    CommonModule,
    NzFormModule,
    CodemirrorModule,
    FormsModule,
    NzUploadModule,
    NzButtonModule,
    NzGridModule,
    ReactiveFormsModule,
    NzNotificationModule,
    NzModalModule
  ],
  templateUrl: './code.component.html',
  styleUrl: './code.component.css'
})
export class CodeComponent {

  constructor(
    private scriptService: ScriptService,
    private fb: NonNullableFormBuilder,
    private notification:NzNotificationService,
    private modal: NzModalService,
  ) {
    this.script = this.fb.group({
      name: ["", [Validators.required]],
      type: ["影视", [Validators.required]],
      languageType: ["go", [Validators.required]],
      content: [`package main

import "fmt"
      
func main(){
  fmt.Println("test")
}`, [Validators.required]],
      cycle: [60, [Validators.min(60)]],
      status: [false],
      description:["",[Validators.required]]
    });
  }

  username:string="ysm"   // 等实现登陆后获取
  scriptTypes = ['影视', '书籍'];  // 脚本类型
  languageTypes = ["go"] // 语言类型
  script: FormGroup<{
    name: FormControl<string>;
    type: FormControl<string>;
    languageType: FormControl<string>;
    content: FormControl<string>;
    cycle: FormControl<number>;
    status: FormControl<boolean>;
    description:FormControl<string>;
  }>;
  statusOptions = [{label:"启用",value:true},{label:"禁用",value:false}];


  // 添加脚本类型
  index = 0;
  addScriptType(input: HTMLInputElement): void {
    const value = input.value
    if (value === "") {
      console.log("脚本类型不能填写为空");
      return;
    }
    if (this.scriptTypes.indexOf(value) === -1) {
      this.scriptTypes = [...this.scriptTypes, value];
    }
  }

  public codeOption = {
    lineNumbers: true,
    lineWrapping: true,
    tabSize: 2,
    theme: 'blackboard',
    mode: "go",          //定义mode
  }

  addScript() {
    var name:string;
    var type:string;
    var code:string;
    var status:boolean;
    var description:string;
    var language:string;
    var cycle:number;
    if(this.script.value.name === undefined || this.script.value.name === ""){
      this.notification.error("验证失败","脚本名字不能为空");
      return
    }else{
      name = this.script.value.name
    }

    if(this.script.value.type === undefined || this.script.value.type===""){
      this.notification.error("验证失败","脚本类型不能为空");
      return
    }else{
      type = this.script.value.type
    }

    if(this.script.value.content === undefined || this.script.value.content === ""){
      this.notification.error("验证失败","代码不能为空");
      return
    }else{
      code = this.script.value.content
    }

    if(this.script.value.status === undefined){
      this.notification.error("验证失败","请检查下运行状态");
      return
    }else{
      status = this.script.value.status
    }

    if(this.script.value.description === undefined || this.script.value.description == ""){
      this.notification.error("验证失败","请输入相关描述");
      return
    }else{
      description = this.script.value.description
    }

    if(this.script.value.languageType === undefined || this.script.value.languageType == ""){
      this.notification.error("验证失败","请输入语言类型");
      return
    }else{
      language = this.script.value.languageType
    }

    if(this.script.value.cycle === undefined || this.script.value.cycle < 60){
      this.notification.error("验证失败","请输入运行周期，必须大于60s");
      return
    }else{
      cycle = this.script.value.cycle
    }
    if (!this.script.valid){
      this.notification.error("验证失败","请检查下相关选项是否正确");
      return 
    }

    var s:script={
      username:this.username,
      name:name,
      type:type,
      code:code,
      status:status,
      description:description,
      language:language,
      cycle:cycle
    }
    this.scriptService.addScript(s).subscribe((r:response)=>{
      if (r.status!==0){
        this.notification.error("新增爬虫脚本失败",r.msg);
        return;
      }
      this.notification.info("新增爬虫脚本成功","");
    });
  }

  runDebug(){
    var codeContent:string;
    if(this.script.value.content === undefined || this.script.value.content === ""){
      this.notification.error("数据异常","代码不能为空");
      return
    }else{
      codeContent = this.script.value.content
    }
    var s:script={
      username:this.username,
      code:codeContent
    };
    this.scriptService.runDebug(s).subscribe((r:response)=>{
      if (r.status!==0){
        this.notification.error("脚本运行失败",r.msg);
        return;
      }
      this.modal.info({
        nzTitle:"脚本运行成功",
        nzContent:r.data
      })
    })
  }
}
