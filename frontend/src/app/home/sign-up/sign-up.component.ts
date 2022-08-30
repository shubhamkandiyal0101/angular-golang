import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators, AbstractControl } from "@angular/forms";
import { HomeService } from "src/app/services/home.service";
import { SharedDataService } from "src/app/services/shared-data.service";
import { ToastrService } from 'ngx-toastr';
import { Router } from "@angular/router";

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.scss']
})
export class SignUpComponent implements OnInit {

  signupForm: FormGroup;
  isSubmitted: boolean = false;

  constructor(private formBuilder: FormBuilder, private _sharedData: SharedDataService, private _homeService: HomeService, private toastr: ToastrService, private router: Router) {
    this.buildSignupForm()
  }

  ngOnInit(): void {
    this._sharedData.showHeader.next(false);
  }

  buildSignupForm() {
    this.signupForm = this.formBuilder.group({
      email:new FormControl("",[Validators.required, Validators.email, Validators.maxLength(50)]),
      password: new FormControl("",[Validators.required, Validators.minLength(8), Validators.maxLength(25)]),
      cpassword: new FormControl("",[Validators.required]),
      full_name: new FormControl("", [Validators.required,Validators.maxLength(50)])
    },{validators:this.matchPassword})


  }

  get formControls(){
    return this.signupForm.controls;
  }

  // submit signup form
  submitSignupForm() {


    // console.log(this.signupForm)

    this.isSubmitted = true;

    if(this.signupForm.invalid) {
      return false;
    }

    let formData = this.signupForm.value;
    delete formData["cpassword"];

    this._homeService.signup(formData).subscribe((res)=>{
      this.toastr.success('Success', 'You have signup successfully!');
      this.router.navigate(["/login"])
    },(err)=>{
      this.toastr.error("Error!","Error while doing signup. Please try Again.")
    })
  }
  // ends here ~ submit signup form

  // custom validation 
  matchPassword(control: AbstractControl) {

    if(control) {

      if(control.get("password").value == control.get("cpassword").value) {
        return null;
      } else {
        return {"passwordMismatch":true}
      }
    }
  }
  // ends here ~ custom validation

}
