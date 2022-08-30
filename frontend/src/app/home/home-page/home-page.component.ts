import { Component, OnInit } from '@angular/core';
import { SharedDataService } from "src/app/services/shared-data.service";
import { HttpClient } from "@angular/common/http";
@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.scss']
})
export class HomePageComponent implements OnInit {

  constructor(private _sharedData: SharedDataService, private http:HttpClient) { }

  ngOnInit(): void {
    this._sharedData.showHeader.next(true);
  }

  getData() {
    this.http.get("https://jsonplaceholder.typicode.com/users")
    .subscribe((res)=>console.log(res))
  }

}
