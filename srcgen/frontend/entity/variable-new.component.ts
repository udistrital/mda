import { Component, OnInit } from '@angular/core';
import { Variable } from '../../models/variable';
import { VariableService } from '../../services/variable.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-variable-new',
  templateUrl: './variable-new.component.html',
  styleUrls: ['./variable-new.component.css']
})
export class VariableNewComponent implements OnInit {

  variable: Variable;
  display = false;
  constructor(private variableService: VariableService, private location: Location) { }

  ngOnInit() {
    this.variable = new Variable();
  }

  guardar(variable: Variable): void {

    this.variableService.create(variable);
    this.display = true;

  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}