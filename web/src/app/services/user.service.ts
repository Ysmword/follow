import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, catchError, map, Observable, tap, throwError } from 'rxjs';
import { user } from '../../class/user';
import { response } from '../../class/response';
import { NzMessageService } from 'ng-zorro-antd/message';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private isLoginApi:string="/api/userExist?username=follow&pwd=follow@123456";
  private isAuthenticated =  new BehaviorSubject<boolean>(false);

  constructor(
    private http:HttpClient,
    private message:NzMessageService,
  ) { 
    // 可以在这里添加代码来检查本地存储中的token，并据此设置isAuthenticated的初始值 
    this.isAuthenticated.next(!!localStorage.getItem("follow_auth_token"));
  }

  private httpOptions ={
    headers:new HttpHeaders({'content-Type':'application/json'})
  }

  private handleError(err:HttpErrorResponse){
    if (err.status===0){
      console.error('An error occurred:', err.error)
    }else{
      console.error(`Backend returned code ${err.status}, body was: `, err.error)
    }
    return throwError(() => new Error('Something bad happened; please try again later.'));
  }

  isLogin(u:user){
    return this.http.post<response>(this.isLoginApi,u,this.httpOptions).pipe(
      catchError(this.handleError),
      tap((r:response)=>{  // 提前接受认证，不影响后面的订阅
        if (r.status!=0){
          this.message.error(r.msg);
        }else{
          localStorage.setItem("follow_auth_token","test");
          this.isAuthenticated.next(true);
        }
      })
    )
  }

  logout():void{
    localStorage.removeItem("follow_auth_token");
    this.isAuthenticated.next(false);
  }

  // 提供一个可观察对象来检查用户是否已登录
  getIsAuthenticated():Observable<boolean>{
      return this.isAuthenticated.asObservable();
  }
}
