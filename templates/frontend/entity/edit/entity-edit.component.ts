import { Component, OnInit } from '@angular/core';
import { Jornada } from '../../models/jornada';
import { JornadaService } from '../../services/jornada.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-jornada-edit',
  templateUrl: './jornada-edit.component.html',
  styleUrls: ['./jornada-edit.component.css']
})
export class JornadaEditComponent implements OnInit {

  jornada: Jornada = new Jornada();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private jornadaService: JornadaService) {

  }

  actualizar(jornada: Jornada): void {
    this.jornadaService.update(jornada).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.jornadaService.getJornada(params['id']))
      .subscribe(jornada => this.jornada = jornada);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}
