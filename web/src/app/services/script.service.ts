import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { script } from '../../class/script';
import { response } from '../../class/response';
import { catchError,throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class ScriptService {

  constructor(
    private http:HttpClient,
  ) { }

  private auApi:string = "/api/auScript?username=follow&pwd=follow@123456";
  private runDebugApi:string = "/api/runDebug?username=follow&pwd=follow@123456"
  private getApi:string="/api/getUserAllScript?username=follow&pwd=follow@123456"
  private delApi:string="/api/delScript?username=follow&pwd=follow@123456"
  private getByIDApi:string="/api/getScriptByID?username=follow&pwd=follow@123456"

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
  
  auScript(s:script){
    return this.http.post<response>(this.auApi,s,this.httpOptions).pipe(
      catchError(this.handleError)
    )
  }

  runDebug(s:script){
    return this.http.post<response>(this.runDebugApi,s,this.httpOptions).pipe(
      catchError(this.handleError)
    ) 
  }

  getAllScript(s:script){
    return this.http.post<response>(this.getApi,s,this.httpOptions).pipe(
      catchError(this.handleError)
    )
  }

  delScript(s:script){
    return this.http.post<response>(this.delApi,s,this.httpOptions).pipe(
      catchError(this.handleError)
    )
  }

  getScriptByID(id:number){
    const s:script={id:id}
    return this.http.post<response>(this.getByIDApi,s,this.httpOptions).pipe(
      catchError(this.handleError)
    )
  }
}
