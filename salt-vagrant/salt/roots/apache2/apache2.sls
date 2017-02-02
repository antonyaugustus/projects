# Install Apache
#apache2:
#  pkg.installed:
#    - name: apache2
#    - file: apache2
# Start the Apache service
#  service.running:
#    - enable: True
#    - require:
#      - pkg: apache2
# Copy the configuration template to /etc/apache2/apache2.conf and render grains values.
#  file.managed:
#    - name: /etc/apache2/apache2.conf
#    - source: salt://apache2/apache2.conf
#    - template: jinja
#    - require:
#      - pkg: apache2
apache:
    pkg.installed:
        - name: httpd
    service.running:
        - name: httpd
# Copy the configuration template to /etc/apache2/apache2.conf and render grains values.
    file.managed:
        - name: /etc/httpd/conf/httpd.conf
        - source: salt://apache2/apache2.conf
        - template: jinja
        - require:
            - pkg: httpd

