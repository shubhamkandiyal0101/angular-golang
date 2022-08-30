import { Component, OnInit } from '@angular/core';
import { SharedDataService } from "src/app/services/shared-data.service";

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.scss']
})
export class FooterComponent implements OnInit {

  showFooter: boolean = true;

  constructor(private _sharedData: SharedDataService) { }
  ngOnInit(): void {
    this._sharedData.showHeader.subscribe((res)=>{
      // console.log(res);
      this.showFooter = res;
    })
  }
  

}
