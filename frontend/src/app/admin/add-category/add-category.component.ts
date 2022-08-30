import { Component, OnInit } from '@angular/core';
import { FormGroup, Validators, FormControl, FormBuilder } from '@angular/forms';
import { AdminService } from "src/app/services/admin.service";
@Component({
  selector: 'app-add-category',
  templateUrl: './add-category.component.html',
  styleUrls: ['./add-category.component.scss']
})
export class AddCategoryComponent implements OnInit {

  formMode: string = "createMode";
  isSubmitted: boolean = true;
  categoryForm: FormGroup;
  selCategory: any = {};
  categoriesData: any[] = [];

  constructor(private formBuilder:FormBuilder, private _adminService: AdminService) {

    this.buildCategoryForm();
    this.getAllCategories();

  }

  ngOnInit(): void {
  }

  // category form
  buildCategoryForm() {
    this.categoryForm = this.formBuilder.group({
      category_name: new FormControl("", [Validators.required]),
      cat_permalink: new FormControl("", [Validators.required])
    })
  }
  // ends here ~ category form

  // add or edit category data
  addOrEditCategory() {

    this.isSubmitted = true;

    if(this.categoryForm.invalid) {
      return false;
    }

    let formData = this.categoryForm.value;
    
    if(this.formMode == "createMode") {

      // in create mode
      this._adminService.addCategory(formData)
      .subscribe((res)=>{
  
      }, (err)=>{
  
      })
      // ends here ~ in create mode

    } else if (this.formMode == "editMode") {

      // in edit mode
      formData.id = this.selCategory.id;
      const payload = formData;
      this._adminService.updateCategory(payload)
      .subscribe((res)=>{
  
      }, (err)=>{
  
      })
      // ends here ~ in edit mode

    }
    
  }
  // ends here ~ add or edit category data

  // get all categories
  getAllCategories() {
    this._adminService.getAllCategories().subscribe((res)=>{
      this.categoriesData = [];
    },(err)=>{

    })

  }
  // ends here ~ get all categories


}
