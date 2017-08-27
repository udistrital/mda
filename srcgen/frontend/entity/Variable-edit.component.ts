import { Component, OnInit } from '@angular/core';
import { Variable } from '../../models/variable';
import { VariableService } from '../../services/variable.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-variable-edit',
  templateUrl: './variable-edit.component.html',
  styleUrls: ['./variable-edit.component.css']
})
export class VariableEditComponent implements OnInit {

  variable: Variable = new Variable();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private variableService: VariableService) {

  }

  actualizar(variable: Variable): void {
    this.variableService.update(variable).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.variableService.getVariable(params['id']))
      .subscribe(variable => this.variable = variable);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}