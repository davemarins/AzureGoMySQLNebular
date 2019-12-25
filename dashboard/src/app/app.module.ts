import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {NbLayoutModule, NbThemeModule} from '@nebular/theme';
import {NbEvaIconsModule} from '@nebular/eva-icons';
// Nebular unique module importer
import {NebularModule} from "./nebular.module";
// Nebular sidebar service
import {NbSidebarService} from "@nebular/theme";

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    NbThemeModule.forRoot({name: 'default'}),
    NbLayoutModule,
    NbEvaIconsModule,
    // Nebular unique module importer
    NebularModule
  ],
  // Nebular sidebar service
  providers: [NbSidebarService],
  bootstrap: [AppComponent]
})
export class AppModule {
}
