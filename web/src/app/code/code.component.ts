import { CommonModule } from '@angular/common';
import { Component,ViewChild } from '@angular/core';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzFormModule } from 'ng-zorro-antd/form';
import { CodemirrorModule } from '@ctrl/ngx-codemirror';
import { FormsModule } from '@angular/forms';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzGridModule } from 'ng-zorro-antd/grid';
import 'codemirror/theme/blackboard.css';
import 'codemirror/addon/hint/sql-hint.js';
import 'codemirror/mode/sql/sql';
import 'codemirror/addon/hint/show-hint.css';
import 'codemirror/addon/hint/show-hint.js';

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
  ],
  templateUrl: './code.component.html',
  styleUrl: './code.component.css'
})
export class CodeComponent {
  codeName?: string; // 脚本名称
  codeContent?:string; // 脚本内容

  listOfItem = ['影视','书籍'];
  index = 0;
  addItem(input:HTMLInputElement):void{
      const value = input.value
      if (this.listOfItem.indexOf(value) === -1) {
        this.listOfItem = [...this.listOfItem, input.value];
      }
  }

  public codeOption = {
    lineNumbers: true,
    lineWrapping: true,
    tabSize: 2,
    theme: 'blackboard',
    mode:  "text/x-mysql",          //定义mode
    extraKeys: {"Ctrl": "autocomplete"},   //自动提示配置
  }
}
