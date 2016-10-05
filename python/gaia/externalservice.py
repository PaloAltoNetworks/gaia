# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class ExternalService(RESTObject):
    """ Represents a ExternalService in the 

        Notes:
            ExternalService are services that are external to the system.
    """

    def __init__(self, **kwargs):
        """ Initializes a ExternalService instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> externalservice = ExternalService(id=u'xxxx-xxx-xxx-xxx', name=u'ExternalService')
              >>> externalservice = ExternalService(data=my_dict)
        """

        super(ExternalService, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._annotation = None
        self._associatedtags = None
        self._createdat = None
        self._deleted = None
        self._description = None
        self._name = None
        self._namespace = None
        self._network = None
        self._parentid = None
        self._parenttype = None
        self._port = None
        self._protocol = None
        self._status = None
        self._updatedat = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="deleted", remote_name="deleted")
        self.expose_attribute(local_name="description", remote_name="description")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="network", remote_name="network")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="port", remote_name="port")
        self.expose_attribute(local_name="protocol", remote_name="protocol")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="updatedAt", remote_name="updatedAt")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        return self.ID

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        self.ID = ID

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return externalserviceIdentity

    # Properties
    @property
    def ID(self):
        """ Get ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        return self._id

    @ID.setter
    def ID(self, value):
        """ Set ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        self._id = value
    
    @property
    def annotation(self):
        """ Get annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        return self._annotation

    @annotation.setter
    def annotation(self, value):
        """ Set annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        self._annotation = value
    
    @property
    def associatedTags(self):
        """ Get associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        return self._associatedtags

    @associatedTags.setter
    def associatedTags(self, value):
        """ Set associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        self._associatedtags = value
    
    @property
    def createdAt(self):
        """ Get createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        return self._createdat

    @createdAt.setter
    def createdAt(self, value):
        """ Set createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        self._createdat = value
    
    @property
    def deleted(self):
        """ Get deleted value.

          Notes:
              Deleted marks if the entity has been deleted.

              
        """
        return self._deleted

    @deleted.setter
    def deleted(self, value):
        """ Set deleted value.

          Notes:
              Deleted marks if the entity has been deleted.

              
        """
        self._deleted = value
    
    @property
    def description(self):
        """ Get description value.

          Notes:
              Description is the description of the object.

              
        """
        return self._description

    @description.setter
    def description(self, value):
        """ Set description value.

          Notes:
              Description is the description of the object.

              
        """
        self._description = value
    
    @property
    def name(self):
        """ Get name value.

          Notes:
              Name is the name of the entity

              
        """
        return self._name

    @name.setter
    def name(self, value):
        """ Set name value.

          Notes:
              Name is the name of the entity

              
        """
        self._name = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        self._namespace = value
    
    @property
    def network(self):
        """ Get network value.

          Notes:
              Network refers to either CIDR or domain name

              
        """
        return self._network

    @network.setter
    def network(self, value):
        """ Set network value.

          Notes:
              Network refers to either CIDR or domain name

              
        """
        self._network = value
    
    @property
    def parentID(self):
        """ Get parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        return self._parentid

    @parentID.setter
    def parentID(self, value):
        """ Set parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        self._parentid = value
    
    @property
    def parentType(self):
        """ Get parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        return self._parenttype

    @parentType.setter
    def parentType(self, value):
        """ Set parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        self._parenttype = value
    
    @property
    def port(self):
        """ Get port value.

          Notes:
              Port refers to network port which could be 100-2000 or * to represent all ports

              
        """
        return self._port

    @port.setter
    def port(self, value):
        """ Set port value.

          Notes:
              Port refers to network port which could be 100-2000 or * to represent all ports

              
        """
        self._port = value
    
    @property
    def protocol(self):
        """ Get protocol value.

          Notes:
              Protocol refers to network protocol like TCP/UDP etc or * to represent all protocols

              
        """
        return self._protocol

    @protocol.setter
    def protocol(self, value):
        """ Set protocol value.

          Notes:
              Protocol refers to network protocol like TCP/UDP etc or * to represent all protocols

              
        """
        self._protocol = value
    
    @property
    def status(self):
        """ Get status value.

          Notes:
              Status of an entity

              
        """
        return self._status

    @status.setter
    def status(self, value):
        """ Set status value.

          Notes:
              Status of an entity

              
        """
        self._status = value
    
    @property
    def updatedAt(self):
        """ Get updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        return self._updatedat

    @updatedAt.setter
    def updatedAt(self, value):
        """ Set updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        self._updatedat = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_required_string("name", self.name)

        if err:
            errors.append(err)

        err = validate_required_string("network", self.network)

        if err:
            errors.append(err)

        err = validate_required_string("port", self.port)

        if err:
            errors.append(err)

        err = validate_required_string("protocol", self.protocol)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # externalserviceIdentity represents the Identity of the object
externalserviceIdentity = {"name": "externalservice", "category": "externalservices", "constructor": ExternalService}