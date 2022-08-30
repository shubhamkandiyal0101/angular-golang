import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { SignUpComponent } from './sign-up/sign-up.component';
import { LoginComponent } from './login/login.component';
import { AdminLoginComponent } from './admin-login/admin-login.component';
import { HomePageComponent } from './home-page/home-page.component';
import { HomeLayoutComponent } from './layout/home-layout/home-layout.component';

const routes: Routes = [
  {
    path: "",
    component: HomeLayoutComponent,
    children: [
     
      {
        path: "login",
        component: LoginComponent
      },
      {
        path: "signup",
        component: SignUpComponent
      },
      {
        path: "admin-login",
        component: AdminLoginComponent
      },
      {
        path:"",
        component: HomePageComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class HomeRoutingModule { }
