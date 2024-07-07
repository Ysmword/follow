import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, throwError } from 'rxjs';
import { user } from '../../class/user';
import { response } from '../../class/response';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(
    private http:HttpClient,
  ) { }

  private isLoginApi:string="/api/userExist?username=follow&pwd=follow@123456";

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
      catchError(this.handleError)
    )
  }
  
}
