import { Component, OnInit } from '@angular/core';
import { SharedDataService } from "src/app/services/shared-data.service";
import { FormBuilder, FormControl, FormGroup, Validators } from "@angular/forms";


import { HomeService } from "src/app/services/home.service";
import { ToastrService } from 'ngx-toastr';
import { Router } from "@angular/router";



@Component({
  selector: 'app-admin-login',
  templateUrl: './admin-login.component.html',
  styleUrls: ['./admin-login.component.scss']
})
export class AdminLoginComponent implements OnInit {

  loginForm: FormGroup;
  isSubmitted: boolean = false;

  constructor(private _sharedData: SharedDataService, private formBuilder: FormBuilder, private _homeService: HomeService, private toastr: ToastrService, private router: Router) { 

    this.buildLoginForm();

  }

  ngOnInit(): void {
    this._sharedData.showHeader.next(false);
    const tokenData = localStorage.getItem("admin_token")
    // console.log(" >> ",tokenData)
    if(tokenData?.length > 2) {
      this.router.navigate(["/admin"])
    }
  }

  // login form
  buildLoginForm() {

    this.loginForm = this.formBuilder.group({
      email: new FormControl("",[Validators.required,Validators.email]),
      password: new FormControl("",[Validators.required, Validators.minLength(8)])
    })

  }
  // ends here ~ login form

  //on submit login form
  submitLoginForm() {
    this.isSubmitted = true;

    if(this.loginForm.invalid) {
      return false;
    }

    let formData = this.loginForm.value;
    formData["is_admin"] = true;

    this._homeService.loginUser(formData)
    .subscribe((res:any)=>{
      const token = res.data.token;
      localStorage.setItem("admin_token",token)
      this.router.navigate(["/admin"]);
      this.toastr.success('Success', 'You have logged in successfully!');

    },(err)=>{
      this.toastr.error('Error', 'Please provide correct credentials to login');

    })

  }
  // ends here ~ on submit login form

}
