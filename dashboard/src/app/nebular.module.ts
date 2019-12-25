import {NgModule} from "@angular/core";
import {NbActionsModule, NbButtonModule, NbCardModule, NbSidebarModule} from "@nebular/theme";

@NgModule({
  imports: [
    NbButtonModule,
    NbCardModule,
    NbActionsModule,
    NbSidebarModule
  ],
  exports: [
    NbButtonModule,
    NbCardModule,
    NbActionsModule,
    NbSidebarModule
  ]
})
export class NebularModule {
}
