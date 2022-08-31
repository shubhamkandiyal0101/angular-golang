import { Component, OnInit } from '@angular/core';
import { FormGroup, Validators, FormControl, FormBuilder } from '@angular/forms';
import { AdminService } from "src/app/services/admin.service";
import { ToastrService } from 'ngx-toastr';

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
  showForm: boolean = false;

  constructor(private formBuilder:FormBuilder, private _adminService: AdminService, private toastr: ToastrService) {

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
        this.toastr.success("Success","New Category Added Successfully");
        window.location.reload();
      }, (err)=>{
        this.toastr.error("Error","Something went wrong while adding category. Please try again")
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
    this._adminService.getAllCategories().subscribe((res:any[])=>{
      this.categoriesData = res;
    },(err)=>{

    })

  }
  // ends here ~ get all categories

  // on click edit category
  onClickEditCat(index) {
    this.formMode = "editMode";
    this.showForm = true;
    this.selCategory = this.categoriesData[index];
    this.categoryForm.get("category_name").patchValue(this.selCategory.category_name);
    this.categoryForm.get("cat_permalink").patchValue(this.selCategory.cat_permalink);
  }
  // ends here ~ on click edit category

  deleteCat(index) {
    let selCatId = this.categoriesData[index]._id;
    console.log(" >> ",this.categoriesData[index])
    this._adminService.deleteProductCat(selCatId)
    .subscribe((res)=>{
      this.toastr.success("Success","Category Deleted Successfully");
      this.categoriesData.splice(index,1);
    })
  }

  cancelCatForm(){
    this.formMode = "createMode";
    this.showForm = false;
  }

  showAddCatForm() {
    this.formMode = "createMode";
    this.categoryForm.reset();
    this.showForm = true;
  }
  


}
