import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { script } from '../../class/script';
import { response } from '../../class/response';

@Injectable({
  providedIn: 'root'
})

export class ScriptService {

  constructor(
    private http:HttpClient,
  ) { }

  private addApi:string = "/api/addScript?username=follow&pwd=follow@123456";

  httpOptions ={
    headers:new HttpHeaders({'content-Type':'application/json'})
  }
  
  addScript(s:script){
    s.id =0 ;
    return this.http.post<response>(this.addApi,s,this.httpOptions) 
  }
}
