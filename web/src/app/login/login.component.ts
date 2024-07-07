import { Component, EventEmitter, Output } from '@angular/core';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { FormControl, FormGroup, FormsModule, NonNullableFormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { user } from '../../class/user';
import { UserService } from '../services/user.service';
import { response } from '../../class/response';
import { NzMessageService } from 'ng-zorro-antd/message';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    NzFormModule,
    NzInputModule,
    FormsModule,
    NzButtonModule,
    ReactiveFormsModule,
    NzCheckboxModule,
    NzIconModule
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {

  userInfo: FormGroup<{
    userName: FormControl<string>;
    password: FormControl<string>;
    remember: FormControl<boolean>;
  }> = this.fb.group({
    userName: ['', [Validators.required]],
    password: ['', [Validators.required]],
    remember: [true]
  });
  @Output() childEvent = new EventEmitter<boolean>();

  submitForm(): void {
    if (this.userInfo.valid) {
      var u:user={
        username:this.userInfo.value.userName,
        password:this.userInfo.value.password
      }
      this.userServce.isLogin(u).subscribe((r:response)=>{
        console.log(u,r)
        if (r.status!=0){
          this.message.error(r.msg);
        }else{
          this.message.info(r.msg);
          this.sendToParent()
        }
      })
    } else {
      Object.values(this.userInfo.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }
 
  sendToParent() {
    this.childEvent.emit(true);
  }


  constructor(
    private fb: NonNullableFormBuilder,
    private userServce:UserService,
    private message:NzMessageService,
    ) {}
}
