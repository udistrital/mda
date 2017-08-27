"""
An example how to generate angularjs code from textX model using jinja2
template engine (http://jinja.pocoo.org/docs/dev/)
"""
from os import mkdir
from os.path import exists, dirname, join
import jinja2
from entity_test import get_entity_mm


def main(debug=False):

    this_folder = dirname(__file__)

    entity_mm = get_entity_mm(debug)

    # Build Person model from person.ent file
    entity_model = entity_mm.model_from_file(join(this_folder, 'voto.ent'))

    #print entity_model.entities

    def is_entity(n):
        """
        Test to prove if some type is an entity
        """
        if n.type in entity_model.entities:
            return True
        else:
            return False

    def beegotype(s):
        """
        Maps type names from PrimitiveType to beego.
        """
        return {
                'integer': 'int',
                'string': 'string',
                'bool': 'boolean',
                'time': 'Time.time'
        }.get(s.name, s.name)

    def angulartype(s):
        """
        Maps type names from PrimitiveType to angular.
        """
        return {
                'integer': 'int',
                'string': 'string',
                'bool': 'boolean',
                'time': 'Date'
        }.get(s.name, s.name)

    # Create output folder for backend
    srcgen_folder = join(this_folder, 'srcgen')
    if not exists(srcgen_folder):
        mkdir(srcgen_folder)
    srcgen_folder_backend = join(this_folder, 'srcgen/backend')
    if not exists(srcgen_folder_backend):
        mkdir(srcgen_folder_backend)
    srcgen_folder_controler = join(this_folder, 'srcgen/backend/controllers')
    if not exists(srcgen_folder_controler):
        mkdir(srcgen_folder_controler)
    srcgen_folder_model = join(this_folder, 'srcgen/backend/models')
    if not exists(srcgen_folder_model):
        mkdir(srcgen_folder_model)
    # Create output folder for frontend
    srcgen_folder_frontend = join(this_folder, 'srcgen/frontend')
    if not exists(srcgen_folder_frontend):
        mkdir(srcgen_folder_frontend)
    srcgen_folder_frontend_model = join(this_folder, 'srcgen/frontend/models')
    if not exists(srcgen_folder_frontend_model):
        mkdir(srcgen_folder_frontend_model)

    # Initialize template engine.
    jinja_env = jinja2.Environment(
        loader=jinja2.FileSystemLoader(this_folder),
        trim_blocks=True,
        lstrip_blocks=True)

    # Register filter for mapping Entity type names to beego type names.
    jinja_env.filters['beegotype'] = beegotype
    # Register filter for mapping Entity type names to angular type names.
    jinja_env.filters['angulartype'] = angulartype

    jinja_env.tests['entity'] = is_entity

    # Load template
    template = jinja_env.get_template('templates/backend/controller.template')

    for entity in entity_model.entities:
        # For each entity generate java file
        with open(join(srcgen_folder_controler, "%s.go" % entity.name.capitalize()), 'w') as f:
            f.write(template.render(entity=entity))

    # Load template
    template = jinja_env.get_template('templates/backend/model.template')

    for entity in entity_model.entities:
        with open(join(srcgen_folder_model, "%s.go" % entity.name.capitalize()), 'w') as f:
            f.write(template.render(entity=entity))

    # Load template
    template = jinja_env.get_template('templates/frontend/models/entity.ts.template')

    for entity in entity_model.entities:
        with open(join(srcgen_folder_frontend_model, "%s.go" % entity.name.capitalize()), 'w') as f:
            f.write(template.render(entity=entity))

if __name__ == "__main__":
    main()
