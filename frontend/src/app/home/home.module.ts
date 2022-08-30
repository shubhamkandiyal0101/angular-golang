import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { HomeRoutingModule } from './home-routing.module';
import { SignUpComponent } from './sign-up/sign-up.component';
import { LoginComponent } from './login/login.component';
import { AdminLoginComponent } from './admin-login/admin-login.component';
import { HomePageComponent } from './home-page/home-page.component';
import { HeaderComponent } from './layout/header/header.component';
import { FooterComponent } from './layout/footer/footer.component';
import { HomeLayoutComponent } from './layout/home-layout/home-layout.component';
import { HttpClientModule, HTTP_INTERCEPTORS } from "@angular/common/http";

import { TokenInterceptor } from "src/app/auth/token.interceptor";

import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { ToastrModule } from 'ngx-toastr';

@NgModule({
  declarations: [
    SignUpComponent,
    LoginComponent,
    AdminLoginComponent,
    HomePageComponent,
    HeaderComponent,
    FooterComponent,
    HomeLayoutComponent
  ],
  imports: [
    CommonModule,
    HomeRoutingModule,
    HttpClientModule,
    FormsModule, 
    ReactiveFormsModule,
    ToastrModule.forRoot(),
  ],
  providers: []
})
export class HomeModule { }
