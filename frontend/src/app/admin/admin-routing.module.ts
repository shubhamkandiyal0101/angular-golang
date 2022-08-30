import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AdminDashboardComponent } from './admin-dashboard/admin-dashboard.component';
import { AdminLayoutComponent } from './admin-layout/admin-layout.component';
import { AddCategoryComponent } from './add-category/add-category.component';


const routes: Routes = [
  {
    path: "",
    component: AdminLayoutComponent,
    children: [
      {
        path:"",
        component: AdminDashboardComponent
      },
      {
        path: "manage-category",
        component: AddCategoryComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AdminRoutingModule { }
