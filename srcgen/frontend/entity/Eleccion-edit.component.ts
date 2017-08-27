import { Component, OnInit } from '@angular/core';
import { Eleccion } from '../../models/eleccion';
import { EleccionService } from '../../services/eleccion.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-eleccion-edit',
  templateUrl: './eleccion-edit.component.html',
  styleUrls: ['./eleccion-edit.component.css']
})
export class EleccionEditComponent implements OnInit {

  eleccion: Eleccion = new Eleccion();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private eleccionService: EleccionService) {

  }

  actualizar(eleccion: Eleccion): void {
    this.eleccionService.update(eleccion).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.eleccionService.getEleccion(params['id']))
      .subscribe(eleccion => this.eleccion = eleccion);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}