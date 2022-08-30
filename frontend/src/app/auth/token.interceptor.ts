import { Injectable } from '@angular/core';
import { Router } from "@angular/router";
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class TokenInterceptor implements HttpInterceptor {

  constructor(private router:Router) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    // console.log(window.location.href);
    let location = this.router.url;
    let splitLocation = location.split("/");

    const userRoute = splitLocation[1]

    if(userRoute == "admin") {
        const API_TOKEN = localStorage.getItem("admin_token")
        // request = request.clone({setHeaders:{"api-key":"API_KEY"}})
        return next.handle(request.clone({setHeaders:{"token":API_TOKEN} }));
    } else {
        return next.handle(request)
    }

    
  }
}
