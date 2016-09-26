# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class Notification(RESTObject):
    """ Represents a Notification in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a Notification instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> notification = Notification(id=u'xxxx-xxx-xxx-xxx', name=u'Notification')
              >>> notification = Notification(data=my_dict)
        """

        super(Notification, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._annotation = None
        self._associatedtags = None
        self._createdat = None
        self._deleted = None
        self._deletedat = None
        self._limit = None
        self._name = None
        self._namespace = None
        self._new = None
        self._nextpage = None
        self._notified = None
        self._old = None
        self._page = None
        self._parentid = None
        self._parenttype = None
        self._status = None
        self._updatedat = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="deleted", remote_name="deleted")
        self.expose_attribute(local_name="deletedAt", remote_name="deletedAt")
        self.expose_attribute(local_name="limit", remote_name="limit")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="new", remote_name="new")
        self.expose_attribute(local_name="nextPage", remote_name="nextPage")
        self.expose_attribute(local_name="notified", remote_name="notified")
        self.expose_attribute(local_name="old", remote_name="old")
        self.expose_attribute(local_name="page", remote_name="page")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
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
        return notificationIdentity

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
    def deletedAt(self):
        """ Get deletedAt value.

          Notes:
              Deleted is the time when the notification was deleted

              
        """
        return self._deletedat

    @deletedAt.setter
    def deletedAt(self, value):
        """ Set deletedAt value.

          Notes:
              Deleted is the time when the notification was deleted

              
        """
        self._deletedat = value
    
    @property
    def limit(self):
        """ Get limit value.

          Notes:
              Limits the amount of results in the "LayersIntroducingVulnerability" property on New and Old vulnerabilities

              
        """
        return self._limit

    @limit.setter
    def limit(self, value):
        """ Set limit value.

          Notes:
              Limits the amount of results in the "LayersIntroducingVulnerability" property on New and Old vulnerabilities

              
        """
        self._limit = value
    
    @property
    def name(self):
        """ Get name value.

          Notes:
              Name is the name of the notification

              
        """
        return self._name

    @name.setter
    def name(self, value):
        """ Set name value.

          Notes:
              Name is the name of the notification

              
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
    def new(self):
        """ Get new value.

          Notes:
              New is the new layers that introduced vulnerability

              
        """
        return self._new

    @new.setter
    def new(self, value):
        """ Set new value.

          Notes:
              New is the new layers that introduced vulnerability

              
        """
        self._new = value
    
    @property
    def nextPage(self):
        """ Get nextPage value.

          Notes:
              NextPage is the next page number

              
        """
        return self._nextpage

    @nextPage.setter
    def nextPage(self, value):
        """ Set nextPage value.

          Notes:
              NextPage is the next page number

              
        """
        self._nextpage = value
    
    @property
    def notified(self):
        """ Get notified value.

          Notes:
              Norified is the time when the notification was sent

              
        """
        return self._notified

    @notified.setter
    def notified(self, value):
        """ Set notified value.

          Notes:
              Norified is the time when the notification was sent

              
        """
        self._notified = value
    
    @property
    def old(self):
        """ Get old value.

          Notes:
              Old is the old layers that introduced vulnerability

              
        """
        return self._old

    @old.setter
    def old(self, value):
        """ Set old value.

          Notes:
              Old is the old layers that introduced vulnerability

              
        """
        self._old = value
    
    @property
    def page(self):
        """ Get page value.

          Notes:
              Page is the page number

              
        """
        return self._page

    @page.setter
    def page(self, value):
        """ Set page value.

          Notes:
              Page is the page number

              
        """
        self._page = value
    
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

        if len(errors) > 0:
            return errors

        return None

    # notificationIdentity represents the Identity of the object
notificationIdentity = {"name": "notification", "category": "notifications", "constructor": Notification}