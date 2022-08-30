import { Component, OnInit } from '@angular/core';
import { SharedDataService } from "src/app/services/shared-data.service";
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  showHeader: boolean = true;

  constructor(private _sharedData: SharedDataService) { }
  ngOnInit(): void {
    this._sharedData.showHeader.subscribe((res)=>{
      // console.log(res);
      this.showHeader = res;
    })
  }
  

}
