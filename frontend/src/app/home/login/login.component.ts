import { Component, OnInit } from '@angular/core';
import { SharedDataService } from "src/app/services/shared-data.service";
import { FormBuilder, FormControl, FormGroup, Validators } from "@angular/forms";


import { HomeService } from "src/app/services/home.service";
import { ToastrService } from 'ngx-toastr';
import { Router } from "@angular/router";



@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;
  isSubmitted: boolean = false;

  constructor(private _sharedData: SharedDataService, private formBuilder: FormBuilder, private _homeService: HomeService, private toastr: ToastrService, private router: Router) { 

    this.buildLoginForm();

  }

  ngOnInit(): void {
    this._sharedData.showHeader.next(false);
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
    formData["is_admin"] = false;

    this._homeService.loginUser(formData)
    .subscribe((res)=>{
      this.router.navigate["/dashboard"];
      this.toastr.success('Success', 'You have logged in successfully!');

    },(err)=>{
      this.toastr.error('Error', 'Please provide correct credentials to login');

    })

  }
  // ends here ~ on submit login form

}
