import { Component, OnInit } from '@angular/core';
import { Ponderacion } from '../../models/ponderacion';
import { PonderacionService } from '../../services/ponderacion.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-ponderacion-edit',
  templateUrl: './ponderacion-edit.component.html',
  styleUrls: ['./ponderacion-edit.component.css']
})
export class PonderacionEditComponent implements OnInit {

  ponderacion: Ponderacion = new Ponderacion();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private ponderacionService: PonderacionService) {

  }

  actualizar(ponderacion: Ponderacion): void {
    this.ponderacionService.update(ponderacion).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.ponderacionService.getPonderacion(params['id']))
      .subscribe(ponderacion => this.ponderacion = ponderacion);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}