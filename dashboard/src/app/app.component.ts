import {Component, OnInit} from '@angular/core';
import {NbSidebarService} from "@nebular/theme";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{

  public title = 'davidemarino.io CMS';
  private subHeader = false;

  constructor(private sidebarService: NbSidebarService) {
  }

  ngOnInit() {
    this.subHeader = false;
    this.sidebarService.toggle(this.subHeader);
  }

  toggle() {
    this.subHeader = !this.subHeader;
    this.sidebarService.toggle(this.subHeader);
  }

}
