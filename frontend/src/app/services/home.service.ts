import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
@Injectable({
  providedIn: 'root'
})
export class HomeService {

  constructor(private http:HttpClient) { }

  signup(data:any) {
    let apiUrl = "/api/user-signup";
    return this.http.post(apiUrl,data);
  }

  loginUser(data:any) {
    let apiUrl = "/api/user-login";
    return this.http.post(apiUrl, data);
  }
}
