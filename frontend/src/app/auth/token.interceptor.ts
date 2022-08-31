import { Injectable } from '@angular/core';
import { Router } from "@angular/router";
import { catchError } from "rxjs/operators"
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable, throwError } from 'rxjs';

@Injectable()
export class TokenInterceptor implements HttpInterceptor {

  constructor(private router:Router) {}

  private handleError(err:any) {
    // console.log(err)


    const errCodeStr = err.status;
    if(`${errCodeStr}`.startsWith("5")) {
      localStorage.clear();
      this.router.navigate(["/admin-login"])
    }
    
    return throwError(err)
  }


  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    // console.log(window.location.href);
    let location = this.router.url;
    let splitLocation = location.split("/");

    const userRoute = splitLocation[1]

    if(userRoute == "admin") {
        const API_TOKEN = localStorage.getItem("admin_token")
        // request = request.clone({setHeaders:{"api-key":"API_KEY"}})
        return next.handle(request.clone({setHeaders:{'Authorization': `Bearer ${API_TOKEN}`
      } })).pipe(catchError(x=> this.handleError(x)));
    } else {
        return next.handle(request)
    }

    
    
  }
}
