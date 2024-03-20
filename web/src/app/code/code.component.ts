import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzFormModule } from 'ng-zorro-antd/form';
import { CodemirrorModule } from '@ctrl/ngx-codemirror';
import { FormsModule } from '@angular/forms';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { NzButtonModule } from 'ng-zorro-antd/button';

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
}
