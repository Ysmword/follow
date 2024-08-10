import { Component } from '@angular/core';
import { RouterModule, RouterOutlet } from '@angular/router';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { LoginComponent } from './login/login.component';
import { CommonModule } from '@angular/common';


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    NzLayoutModule,
    NzButtonModule,
    NzMenuModule,
    NzIconModule,
    RouterModule,
    LoginComponent,
    CommonModule
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'follow';
  isLogined:boolean=false;

  receiveChildEvent(event: boolean) {
    this.isLogined = event;
    localStorage.setItem('isLoggedIn', 'true');
    localStorage.setItem('loginTime', new Date().getTime().toString());
  }

  constructor(
  ){
    const isLoggedIn = localStorage.getItem('isLoggedIn');
    const loginTime = parseInt(localStorage.getItem('loginTime') || '0', 10);
    const currentTime = new Date().getTime();
    this.isLogined= isLoggedIn === 'true' && currentTime - loginTime < 600;
  }
}
