import { HttpClient,HttpErrorResponse,HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError,throwError } from 'rxjs';
import { response } from '../../class/response';
import { result } from '../../class/result';

@Injectable({
  providedIn: 'root'
})
export class ResultService {

  constructor(
    private http:HttpClient,
  ) { }


  private getAllResultApi:string = "/api/getAllResult?username=follow&pwd=follow@123456"
  private getResultByUApi:string = "/api/getResultByU?username=follow&pwd=follow@123456&page="
  private getResultByUTApi:string= "/api/getResultByUT?username=follow&pwd=follow@123456"

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

  getAllResult(){
    return this.http.get<response>(this.getAllResultApi).pipe(
      catchError(this.handleError)
    )
  }

  getResultByU(username:string,page:number){
    var api = this.getResultByUApi + page.toString()
    const r:result={username:username}
    return this.http.post<response>(api,r,this.httpOptions).pipe(
      catchError(this.handleError)
    )
  }

  getResultByUT(username:string,type:string){
    const r:result={username:username,type:type}
    return this.http.post<response>(this.getResultByUTApi,r,this.httpOptions).pipe(
      catchError(this.handleError)
    )
  }
}
