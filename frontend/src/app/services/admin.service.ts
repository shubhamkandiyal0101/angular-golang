import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
@Injectable({
  providedIn: 'root'
})
export class AdminService {

  constructor(private http:HttpClient) { }

  getAllCategories() {
    let apiUrl = "/api/get-all-categories";
    return this.http.get(apiUrl);
  }

  addCategory(payload) {
    let apiUrl = "/api/add-product-category";
    return this.http.post(apiUrl, payload);
  }

  updateCategory(payload) {
    let apiUrl = "/api/";
    return this.http.post(apiUrl, payload);
  }

  deleteProductCat(cat_id) {
    let apiUrl = `/api/delete-product-category/${cat_id}`;
    return this.http.delete(apiUrl)
  }
  
}
