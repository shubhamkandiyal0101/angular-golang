import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AdminRoutingModule } from './admin-routing.module';
import { AdminDashboardComponent } from './admin-dashboard/admin-dashboard.component';
import { AddCategoryComponent } from './add-category/add-category.component';
import { AddProductsComponent } from './add-products/add-products.component';
import { AdminLayoutComponent } from './admin-layout/admin-layout.component';

import { AdminSidenavComponent } from './admin-sidenav/admin-sidenav.component';

import { ReactiveFormsModule, FormsModule } from "@angular/forms";
import { HttpClientModule, HTTP_INTERCEPTORS } from "@angular/common/http";
import { TokenInterceptor } from "src/app/auth/token.interceptor";

@NgModule({
  declarations: [
    AdminDashboardComponent,
    AddCategoryComponent,
    AddProductsComponent,
    AdminLayoutComponent,
    AdminSidenavComponent
  ],
  imports: [
    CommonModule,
    AdminRoutingModule,
    ReactiveFormsModule,
    FormsModule
  ],
  providers: [
   
  ]
})
export class AdminModule { }
